package easyv5

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/767829413/fybq/util"
)

const (
	reward = 50.0
)

type TXInput struct {
	QTXID     []byte // 引用的交易ID
	Index     int64  // 引用的Output的索引值
	Signature []byte // 签名
	PubKey    []byte // 公钥,非原始数据
}

type TXOutput struct {
	Value      float64 // 转账金额
	PubKeyHash []byte  // 收款方公钥哈希
}

// 处理地址到公钥哈希
func (txo *TXOutput) lock(addr string) {
	txo.PubKeyHash = util.GetPubKeyHashByAddr(addr)
}

func NewTXOutput(value float64, addr string) *TXOutput {
	out := &TXOutput{
		Value: value,
	}
	out.lock(addr)
	return out
}

type Transaction struct {
	TXID      []byte      // 交易ID
	TXInputs  []*TXInput  // 交易输入
	TXOutputs []*TXOutput // 交易输出
}

func (tx *Transaction) SetTXID() {
	// 生成交易ID
	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panicln("SetQTXID error: ", err)
	}
	hash := sha256.Sum256(buf.Bytes())
	tx.TXID = hash[:]
}

// 判断是否为挖矿
// 1. Inputs数量为1 2. Input交易id为空 3. Inputs的index为-1
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.TXInputs) == 1 && tx.TXInputs[0].Index == -1 && tx.TXInputs[0].QTXID == nil
}

// 签名
// 参数是私钥和inputs里所有引用的交易结构map[string]Transaction,id作为key
// 1. 创建当前交易副本,将inputs里的 Signature PubKey 置 nil
// 2. 循环遍历副本交易中的inputs,得到每个input对应的公钥hash
// 3. 生成签名需要的数据,也就是对应的hash值
// 3.1 对每一个input进行签名,签名数据为input引用output的hash+当前output
// 3.2 对副本的交易进行hash处理,处理后的hash就是签名的数据
// 4. 进行签名,签名数据放到对应的input Signature字段
func (tx *Transaction) Signature(priKey *ecdsa.PrivateKey, prevTXs map[string]*Transaction) error {
	if tx.IsCoinbase() {
		return nil
	}
	txCopy := tx.TrimCopy()
	for idx, in := range txCopy.TXInputs {
		prevTx, ok := prevTXs[string(in.QTXID)]
		if !ok {
			return fmt.Errorf("signature quoted transaction is null, data is not legal,TXID: %x", in.QTXID)
		}
		// 临时存储公钥hash
		txCopy.TXInputs[idx].PubKey = prevTx.TXOutputs[in.Index].PubKeyHash
		// 获取签名所要的hash数据
		txCopy.SetTXID()
		signHashData := txCopy.TXID
		// 还原数据,防止影响下一条input
		txCopy.TXInputs[idx].PubKey = nil
		sig, err := ecdsa.SignASN1(rand.Reader, priKey, signHashData)
		if err != nil {
			return err
		}
		tx.TXInputs[idx].Signature = sig
	}
	return nil
}

// 校验
// 需要公钥和 txCopy(生成hash数据) 签名
// 1. 得到签名数据和公钥
// 2. 得到签名
// 3. 校验签名
func (tx *Transaction) Verify(prevTXs map[string]*Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}
	txCopy := tx.TrimCopy()

	for idx, in := range tx.TXInputs {
		prevTx, ok := prevTXs[string(in.QTXID)]
		if !ok {
			log.Printf("verify quoted transaction is null, data is not legal,TXID: %x", in.QTXID)
			return false
		}
		// 临时存储公钥hash
		txCopy.TXInputs[idx].PubKey = prevTx.TXOutputs[in.Index].PubKeyHash
		// 获取签名所要的hash数据
		txCopy.SetTXID()
		signHashData := txCopy.TXID
		// 还原数据,防止影响下一条input
		txCopy.TXInputs[idx].PubKey = nil
		// 进行校验
		pubKey, err := util.ParseEccPubKeyBytes(in.PubKey)
		if err != nil {
			log.Printf("verify ParseEccPubKeyBytes,TXID: %x,error: %x", in.QTXID, err)
			return false
		}
		if !ecdsa.VerifyASN1(pubKey, signHashData, in.Signature) {
			log.Printf("verify TXID: %x, illegal inputs", in.QTXID)
			return false
		}
	}
	return true
}

func (tx *Transaction) TrimCopy() Transaction {
	var (
		inputs  []*TXInput
		outputs []*TXOutput
	)
	for _, input := range tx.TXInputs {
		inputs = append(inputs, &TXInput{QTXID: input.QTXID, Index: input.Index, Signature: nil, PubKey: nil})
	}
	for _, output := range tx.TXOutputs {
		outputs = append(outputs, &TXOutput{Value: output.Value, PubKeyHash: output.PubKeyHash})
	}
	return Transaction{TXID: tx.TXID, TXInputs: inputs, TXOutputs: outputs}
}

// 创建Coinbase交易(挖矿交易)
// 挖矿交易特点:
// 只有一个input
// 没有引用交易ID
// 没有引用的Output的索引值
// 矿工由于挖矿时无需指定签名,所以签名sig字段自由填写,一般是矿池名字
func NewCoinbase(addr, data string) *Transaction {
	input := &TXInput{Signature: nil, Index: -1, QTXID: nil, PubKey: []byte(data)}
	output := NewTXOutput(reward, addr)
	tx := &Transaction{
		TXInputs:  []*TXInput{input},
		TXOutputs: []*TXOutput{output},
	}
	tx.SetTXID()
	return tx
}

// 创建普通的转账交易
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	var (
		inputs  []*TXInput
		outputs []*TXOutput
	)
	// 创建交易前需要用钱包,通过地址找到对应的钱包
	wallet, ok := GetWallets().WalletsMap[from]
	if !ok {
		fmt.Println("Sender's wallet not found")
		return nil
	}
	// 1. 找到合理的UTXO集合
	pubKeyHash, err := util.GetPubKeyHash(wallet.PubKey)
	if err != nil {
		fmt.Println("Get public hash key error: ", err.Error())
		return nil
	}
	utxos, calc := bc.FindTransactionUTXOs(pubKeyHash, amount)
	if calc < amount {
		fmt.Printf("%x Current balance is insufficient!", from)
		return nil
	}
	// 2. 将这些utxo转成input
	for idx, idxArr := range utxos {
		for _, index := range idxArr {
			input := &TXInput{
				QTXID:     []byte(idx),
				Index:     index,
				Signature: nil,
				PubKey:    wallet.PubKey,
			}
			inputs = append(inputs, input)
		}
	}
	// 3. 创建output
	output := NewTXOutput(amount, to)
	outputs = append(outputs, output)
	if calc > amount {
		outputs = append(outputs, NewTXOutput(calc-amount, from))
	}
	// 4. 如果有零钱需要找零
	tx := &Transaction{
		TXInputs:  inputs,
		TXOutputs: outputs,
	}
	tx.SetTXID()
	// 做签名
	err = bc.SignatureTXs(wallet.PriKey, tx)
	if err != nil {
		fmt.Println("NewTransaction.SignatureTXs error: ", err.Error())
		return nil
	}
	return tx
}

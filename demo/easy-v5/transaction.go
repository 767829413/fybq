package easyv5

import (
	"bytes"
	"crypto/ecdsa"
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
func (tx *Transaction) Signature(priKey *ecdsa.PrivateKey, prevTXs map[string]Transaction) {}

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
	// 1. 根据inputs进行遍历可以获得对应的TXID
	// 2. 目标交易根据TXID来对应
	// 3. 添加到prevTXs
	prevTXs := make(map[string]Transaction)
	for _, inps := range tx.TXInputs {
		qtx := FindTransactionByID(inps.QTXID)
		prevTXs[string(qtx.TXID)] = qtx
	}

	if !ok {
		fmt.Println("NewTransaction.ParseEccPriKeyBytes is error: ", err)
		return nil
	}
	priKey, err := util.ParseEccPriKeyBytes(wallet.PriKey)
	tx.Signature(priKey, prevTXs)
	return tx
}

func FindTransactionByID(TXID []byte) Transaction {
	return Transaction{}
}

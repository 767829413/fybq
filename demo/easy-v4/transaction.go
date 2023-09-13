package easyv4

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const (
	reward = 12.5
)

type TXInput struct {
	QTXID []byte // 引用的交易ID
	Index int64  // 引用的Output的索引值
	Sig   string // 解锁脚本
}

type TXOutput struct {
	Value      float64 // 转账金额
	PubKeyHash string  // 锁定脚本,用地址模拟
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

// 创建Coinbase交易(挖矿交易)
// 挖矿交易特点:
// 只有一个input
// 没有引用交易ID
// 没有引用的Output的索引值
// 矿工由于挖矿时无需指定签名,所以签名sig字段自由填写,一般是矿池名字
func NewCoinbase(addr, data string) *Transaction {
	input := &TXInput{Sig: data, Index: -1, QTXID: nil}
	output := &TXOutput{Value: reward, PubKeyHash: addr}
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
	// 1. 找到合理的UTXO集合
	utxos, calc := bc.FindTransactionUTXOs(from, amount)
	if calc < amount {
		fmt.Printf("%x Current balance is insufficient!", from)
		return nil
	}
	// 2. 将这些utxo转成input
	for idx, idxArr := range utxos {
		for _, index := range idxArr {
			input := &TXInput{
				QTXID: []byte(idx),
				Index: index,
				Sig:   from,
			}
			inputs = append(inputs, input)
		}
	}
	// 3. 创建output
	output := &TXOutput{Value: amount, PubKeyHash: to}
	outputs = append(outputs, output)
	if calc > amount {
		outputs = append(outputs, &TXOutput{Value: calc - amount, PubKeyHash: from})
	}
	// 4. 如果有零钱需要找零
	tx := &Transaction{
		TXInputs:  inputs,
		TXOutputs: outputs,
	}
	tx.SetTXID()
	return tx
}

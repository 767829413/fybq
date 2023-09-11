package easyv4

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type TXInput struct {
	QTXID []byte // 引用的交易ID
	index int64  // 引用的Output的索引值
	Sig   string // 解锁脚本
}

type TXOutput struct {
	value      float64 // 转账金额
	PubKeyHash string  // 锁定脚本,用地址模拟
}

type Transaction struct {
	TXID      []byte      // 交易ID
	TXInputs  []*TXInput  // 交易输入
	TXOutputs []*TXOutput // 交易输出
}

func (tx *Transaction) SetQTXID() {
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

package easyv5

import (
	"crypto/elliptic"
	"log"

	"github.com/767829413/fybq/util"
	"github.com/btcsuite/btcd/btcutil/base58"
)

type Wallet struct {
	// 私钥
	PriKey []byte
	// 公钥
	PubKey []byte
}

func (w *Wallet) NewAddr() (string, error) {
	pubKey := w.PubKey
	rip160hashValue, err := util.GetPubKeyHash(pubKey)
	if err != nil {
		return "", err
	}
	// 拼接version
	version := byte(00)
	payLoad := append([]byte{version}, rip160hashValue...)
	//checksum
	// 返回前4字节作为校验码
	checkCode := util.Checksum(payLoad)
	payLoad = append(payLoad, checkCode...)
	// 使用btcd package
	return base58.Encode(payLoad), nil

}

func NewWallet() *Wallet {
	priKeyBytes, pubKeyBytes, err := util.GenerateEccKeyBytes(elliptic.P256())
	if err != nil {
		log.Panic(err)
	}
	// 生成公钥
	return &Wallet{
		PriKey: priKeyBytes,
		PubKey: pubKeyBytes,
	}
}

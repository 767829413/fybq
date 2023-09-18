package easyv5

import (
	"crypto/elliptic"
	"crypto/sha256"
	"log"

	"github.com/767829413/fybq/util"
	"github.com/btcsuite/btcd/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	// 私钥
	PriKey []byte
	// 公钥
	PubKey []byte
}

func (w *Wallet) NewAddr() (string, error) {
	pubKey := w.PubKey
	hash := sha256.Sum256(pubKey)
	// 获取rip160hash
	rip160hashValue, err := GetRipemd160Hash(hash[:])
	if err != nil {
		return "", err
	}
	// 拼接version
	version := byte(00)
	payLoad := append([]byte{version}, rip160hashValue...)
	//checksum
	// 返回前4字节作为校验码
	checkCode := checksum(payLoad)
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

func GetRipemd160Hash(data []byte) ([]byte, error) {
	rip160hasher := ripemd160.New()
	_, err := rip160hasher.Write(data)
	if err != nil {
		return nil, err
	}
	return rip160hasher.Sum(nil), nil
}

func checksum(data []byte) []byte {
	// 两次 sha256
	preHash1 := sha256.Sum256(data)
	preHash2 := sha256.Sum256(preHash1[:])
	// 返回前4字节作为校验码
	return preHash2[:4]
}

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	fff := []int{1,2,3,4,5}
	fmt.Println(fff[len(fff)-4:])
	// 创建曲线
	curve := elliptic.P256()
	// 生成私钥]
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	// 生成公钥
	pubKey := privateKey.PublicKey

	data := "sdsdwwewe222222"
	hash := sha256.Sum256([]byte(data))
	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic(err)
	}

	log.Println(ecdsa.VerifyASN1(&pubKey, hash[:], sig))
}

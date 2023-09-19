package easyv5

import (
	"log"
	"os"
	"sync"

	"github.com/767829413/fybq/util"
)

var once sync.Once

const walletsFileName = "wallet.dat"

type Wallets struct {
	WalletsMap map[string]*Wallet
}

func (w *Wallets) CreateWallet() string {
	wallet := NewWallet()
	addr, err := wallet.NewAddr()
	if err != nil {
		log.Panic("CreateWallet", err)
	}
	w.WalletsMap[addr] = wallet
	w.saveToFile()
	return addr
}

func (w *Wallets) saveToFile() {
	x := util.StructToBytes(w)
	err := os.WriteFile(walletsFileName, x, 0600)
	if err != nil {
		log.Panic("saveToFile", err)
	}
}

func (w *Wallets) GetAllAddresses() []string {
	var addresses []string
	for addr := range w.WalletsMap {
		addresses = append(addresses, addr)
	}
	return addresses
}

func GetWallets() *Wallets {
	once.Do(func() {
		_, err := os.Stat(walletsFileName)
		if os.IsNotExist(err) {
			w := &Wallets{
				WalletsMap: make(map[string]*Wallet),
			}
			w.saveToFile()
		}
	})
	// loadFile
	wallets := loadFileByWallets()
	return wallets
}

func loadFileByWallets() *Wallets {
	var wallets Wallets
	x, err := os.ReadFile(walletsFileName)
	if err != nil {
		log.Panic("loadFile", err)
	}
	util.BytesToStruct(x, &wallets)
	return &wallets
}

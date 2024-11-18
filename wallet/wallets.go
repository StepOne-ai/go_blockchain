package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"os"
)

const walletFile = "./tmp/wallets.data"

type Wallets struct {
	Wallets map[string]*Wallet
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws *Wallets) AddWalet() string {
	wallet := MakeWallet()
	address := fmt.Sprintf("%s", wallet.Address())

	ws.Wallets[address] = wallet

	return address
}

func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFile()

	return &wallets, err
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) SaveFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	enc := gob.NewEncoder(&content)
	err := enc.Encode(ws)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets

	content, err := os.ReadFile(walletFile)
	if err != nil {
		return err
	}

	gob.Register(elliptic.P256())
	dec := gob.NewDecoder(bytes.NewReader(content))
	err = dec.Decode(&wallets)
	if err != nil {
		return err
	}

	ws.Wallets = wallets.Wallets
	return nil
}

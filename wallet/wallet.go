package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

const (
	checkSumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PublicKey  []byte
	PrivateKey ecdsa.PrivateKey
}

func (w Wallet) Address() []byte {
	pubKeyHash := PublicKeyHash(w.PublicKey)

	versionedHash := append([]byte{version}, pubKeyHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)

	fmt.Printf("Address: %s\n", address)
	fmt.Printf("Public Key: %x\n", w.PublicKey)
	fmt.Printf("Public Hash: %x\n", pubKeyHash)

	return address
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	pub := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return *privateKey, pub
}

func MakeWallet() *Wallet {
	privateKey, publicKey := NewKeyPair()

	return &Wallet{publicKey, privateKey}
}

func PublicKeyHash(pubKey []byte) []byte {
	hash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(hash[:])
	if err != nil {
		panic(err)
	}

	return hasher.Sum(nil)
}

func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checkSumLength]
}

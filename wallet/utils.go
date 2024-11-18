package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

func Base58Encode(data []byte) []byte {
	enc := base58.Encode(data)

	return []byte(enc)
}

func Base58Decode(enc []byte) []byte {
	dec, err := base58.Decode(string(enc[:]))
	if err != nil {
		log.Panic(err)
	}

	return dec
}

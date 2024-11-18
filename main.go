package main

import (
	. "go_blockchain/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cli := CommandLine{}
	cli.Run()
	// wallet := MakeWallet()
	// wallet.Address()
}

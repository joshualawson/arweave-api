package main

import (
	"fmt"
	"github.com/joshualawson/arweave-api"
	"os"
)

func main() {
	a := arweave.New()

	r, err := a.WalletBalance("dRFuVE-s6-TgmykU4Zqn246AR2PIsf3HhBhZ0t5-WXE")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response: \n%v\n", r)
}

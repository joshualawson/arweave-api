package main

import (
	"fmt"
	"github.com/joshualawson/arweave-api"
	"os"
)

func main() {
	a := arweave.New()

	r, err := a.PeerList()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response: \n%v\n", r)
}

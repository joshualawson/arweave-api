package main

import (
	"fmt"
	"github.com/joshualawson/arweave-api"
	"os"
)

func main() {
	a := arweave.New()

	r, err := a.Block("YuTyalVBTNB9t5KhuRezcIgxVz9PbQsbrcY4Tpkiu8XBPgglGM_Yql5qZd0c9PVG")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response: \n%v\n", r)
}

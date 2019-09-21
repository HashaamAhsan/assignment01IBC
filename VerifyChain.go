package assignment01IBC

import (
	"bytes"
	"fmt"
)

func VerifyChain(head *Block) {
	var PrevBlockHash []byte
	//var Hash []byte

	PrevBlockHash = head.PrevBlockHash
	//Hash = head.Hash
	head = head.Prevpointer

	for true {
		if head.Prevpointer == nil {
			println("Blockchain is not Changed\n")
			break
		}

		res := bytes.Compare(head.Hash, PrevBlockHash)
		if res != 0 {
			fmt.Println("Following Block is Changed")
			fmt.Printf("Data: %s\n\n", head.Data)
			break
		}

		PrevBlockHash = head.Hash
		head = head.Prevpointer
	}
}

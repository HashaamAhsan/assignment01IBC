package assignment01IBC

import "fmt"

func ListBlocks(chainHead *Block) {

	for true {
		fmt.Printf("Prev. hash: %x\n", chainHead.PrevBlockHash)
		fmt.Printf("Data: %s\n", chainHead.Data)
		fmt.Printf("Hash: %x\n", chainHead.Hash)
		fmt.Println()
		if chainHead.Prevpointer != nil {
			chainHead = chainHead.Prevpointer
		} else {
			break
		}
	}
}

package assignment01IBC

import "time"

func NewBlock(data string, prevBlockHash []byte, prevblockpointer *Block) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, prevblockpointer, []byte{}}
	block.SetHash()
	return block
}

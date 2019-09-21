package assignment01IBC

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}, nil)
}

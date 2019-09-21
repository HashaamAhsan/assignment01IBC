package assignment01IBC

func InsertBlock(data string, head *Block) *Block {
	if head == nil {
		head = NewBlock(data, nil, nil)
		//prevBlock := bc.Blocks[len(bc.Blocks)-1]
		//newBlock := NewBlock(data, prevBlock.Hash)
		//bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		newBlock := NewBlock(data, head.Hash, head)
		head = newBlock
	}
	return head
}

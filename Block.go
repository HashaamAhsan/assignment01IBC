package assignment01IBC

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Prevpointer   *Block
	Hash          []byte
}

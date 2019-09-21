package assignment01IBC

import "fmt"

func ChangeBlock(data1 string, data2 string, head *Block) {
	for true {
		if string(head.Data) == data1 {
			head.Data = []byte(data2)
			head.SetHash()
			break
		}
		if head.Prevpointer == nil {
			fmt.Println("The Data not found!")
			break
		} else {
			head = head.Prevpointer
		}
	}
}

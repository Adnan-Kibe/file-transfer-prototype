package main

import (
	"fmt"
	"github/adnan-kibe/file-transfer-prototype/documents"
)

func main() {
	err := documents.FullWalkDown("D:/Golang/file transfer/prototype")
	if err != nil {
		fmt.Println(err)
	}
}
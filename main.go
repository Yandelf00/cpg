package main

import (
	"log"
	"os"

	"github.com/Yandelf00/cpg/setOne"
)


func main() {
	dat, err := os.ReadFile("breakRKXORinpt.txt")
	if err != nil {
		log.Fatal(err)
	}
	decodedDat, err := setOne.Base64Decode(dat)
	if err != nil {
		log.Fatal(err)
	}

	setOne.BreakRKXor(decodedDat)
}

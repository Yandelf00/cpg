package main

import (
	"fmt"
	"log"

	"github.com/Yandelf00/cpg/settwo"
)

// func challsix(){
// 	dat, err := os.ReadFile("breakRKXORinpt.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	decodedDat, err := setOne.Base64Decode(dat)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	setOne.BreakRKXor(decodedDat)
// }

func main() {
	res, err := settwo.PKCSPadding("YELLOW SUBMARINE", 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", res)
}

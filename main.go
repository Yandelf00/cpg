package main

import (
	"fmt"
	"log"

	"github.com/Yandelf00/cpg/setOne"
)

func main() {
	dest, err := setOne.HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dest))
}

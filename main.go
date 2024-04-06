package main

import (
	"cryptopals/hextobase"
	"fmt"
)

func main() {
	todecode := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	decoded, err := hextobase.DecodeHex(todecode)
	if err != nil {
		fmt.Println("error decoding")
		return
	}
	base_encoded := hextobase.EncodeBase(decoded)
	fmt.Printf("%s", base_encoded)
}

package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	todecode := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	slice_to_decode := []byte(todecode)
	dst := make([]byte, hex.DecodedLen(len(slice_to_decode)))
	_, err := hex.Decode(dst, slice_to_decode)
	if err != nil {
		return
	}
	fmt.Printf("%s", dst)
}

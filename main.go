package main

import (
	"cryptopals/hextobase"
	"cryptopals/xor"
	"fmt"
)

func main() {
	todecode := []byte("1c0111001f010100061a024b53535009181c")
	todecode_two := []byte("686974207468652062756c6c277320657965")
	decoded, err := hextobase.DecodeHex(todecode)
	if err != nil {
		fmt.Println("error decoding")
		return
	}
	decoded_two, err := hextobase.DecodeHex(todecode_two)
	if err != nil {
		fmt.Println("error decoding")
		return
	}
	xored_byte, err := xor.Xor_bytes(decoded, decoded_two)
	if err != nil {
		fmt.Println(err)
		return
	}
	encoded_xored_byte := hextobase.EncodeHex(xored_byte)
	fmt.Printf("%s", encoded_xored_byte)
}

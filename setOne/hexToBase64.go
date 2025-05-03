package setOne

import (
	"encoding/base64"
	"encoding/hex"
)

//takes an input string thats encoded in hex 
//and converts it to base64 and returns it as bytes
func HexToBase64(input string) ([]byte, error) {
	src := []byte(input) //transforms the input string to byte
	dst := make([]byte, hex.DecodedLen(len(src))) //makes a byte thats the size of the decoded length
	n, err := hex.Decode(dst, src)//decodes the hex
	if err != nil {
		return nil, err
	}
	data := []byte(dst[:n])
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dest, data)//encodes to base64
	return dest, nil
}

//takes a slice of bytes as input and encodes it 
//to Base64 
func Base64Encode(inpt []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(inpt)))
	base64.StdEncoding.Encode(dst, inpt)
	return dst
}

//takes a slice of bytes encoded in base64
//and decodes it (use this in the challenge 6 of set one)
func Base64Decode(inpt []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(inpt)))
	n, err := base64.StdEncoding.Decode(dst, inpt)
	if err != nil {
		return nil, err
	}

	dst = dst[:n]

	return dst, nil
}

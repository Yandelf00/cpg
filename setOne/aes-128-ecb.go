package setOne

import (
	"crypto/aes"
	"fmt"
	"slices"
)

func AesEcbDecrypt(inpt, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	rest := make([]byte, len(inpt))
	for i, j := 0, 16; i < len(inpt); i, j = i+16, j+16 {
		block.Decrypt(rest[i:j], inpt[i:j])
	}
	return rest, nil
}

func DetectAes(inpt []byte, lnm int) {
	line := 0
	for i, j := 0, 16; i < len(inpt); i, j = i+16, j+16 {
		for k, l := i+16, j+16; k < len(inpt); k, l = k+16, l+16 {
			if slices.Equal(inpt[i:j], inpt[k:l]) {
				line = lnm
			}
		}
	}
	fmt.Println(line)
}

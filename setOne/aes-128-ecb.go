package setOne

import "crypto/aes"

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

package setOne

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(input string) ([]byte, error) {
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	data := []byte(dst[:n])
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dest, data)
	return dest, nil
}

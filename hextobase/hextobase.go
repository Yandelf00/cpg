package hextobase

import (
	"encoding/base64"
	"encoding/hex"
)

func DecodeHex(todecode []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(todecode)))
	_, err := hex.Decode(dst, todecode)
	if err != nil {
		return nil, err
	}
	return dst, err
}

func EncodeBase(toencode []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(toencode)))
	base64.StdEncoding.Encode(eb, toencode)

	return eb
}

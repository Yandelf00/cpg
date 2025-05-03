package setOne

import "encoding/hex"

// takes an input string in hexadecimal and decodes it
func HexDecode(input string) ([]byte, error) {
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

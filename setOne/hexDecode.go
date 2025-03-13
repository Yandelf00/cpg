package setOne

import "encoding/hex"

func HexDecode(input string) ([]byte, error) {
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

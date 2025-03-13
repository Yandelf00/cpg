package setOne

func FixedXor(xoredOne, xoredTwo string) ([]byte, error) {
	inptOne, err := HexDecode(xoredOne)
	if err != nil {
		return nil, err
	}
	inptTwo, err := HexDecode(xoredTwo)
	if err != nil {
		return nil, err
	}
	n := len(inptOne)
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = inptOne[i] ^ inptTwo[i]
	}
	return b, nil
}

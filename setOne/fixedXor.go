package setOne

import "errors"

// takes two equal-length buffers and produces their XOR combination.
func FixedXor(xoredOne, xoredTwo string) ([]byte, error) {
	inptOne, err := HexDecode(xoredOne) //decodes the first xored input
	if err != nil {
		return nil, err
	}
	inptTwo, err := HexDecode(xoredTwo) //decodes the second xored input
	if err != nil {
		return nil, err
	}
	if len(inptOne) != len(inptTwo) {
		return nil, errors.New("The buffers need to have equal-length.")
	}
	n := len(inptOne)
	b := make([]byte, n) //makes a byte to store the result
	for i := 0; i < n; i++ {
		b[i] = inptOne[i] ^ inptTwo[i] //xores the elements of the two inputs
	}
	return b, nil
}

// takes a hex encoded string, decodes it and xores it
// against a single given key
func FixedXorKey(xoredOne string, key byte) ([]byte, error) {
	inptOne, err := HexDecode(xoredOne) //decodes the hex encoded string
	if err != nil {
		return nil, err
	}
	n := len(inptOne)
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = inptOne[i] ^ key //xores every element with the key
	}
	return b, nil
}

// takes a slice of bytes and xores it against a key
func FixedXorKeyByte(xoredOne []byte, key byte) ([]byte, error) {
	n := len(xoredOne)
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = xoredOne[i] ^ key
	}
	return b, nil
}

package xor

import "errors"

func Xor_bytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("the bytes are not of the same length")
	}
	ret := make([]byte, len(a))
	for i := 0; i < len(ret); i++ {
		ret[i] = a[i] ^ b[i]
	}
	return ret, nil
}

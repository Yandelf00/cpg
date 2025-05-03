package setOne

// RepeatingXorKey encrypts (xor encryption) an input with a repeating key
//
// Parameters :
// -input : this is the input string that we are going to encrypt
// -key : the key that we are going to use, key passed as string
//
// Returns :
// -result : the encrypted result with the key, returned as a slice of byte
func RepeatingXorKey(input, key string) []byte {
	var result []byte
	b_inpt := []byte(input)
	b_key := []byte(key)
	n := len(key)
	for r, el := range b_inpt {
		rest := r % n
		if rest == 3 {
			rest = 0
		}
		result = append(result, el^b_key[rest])
	}
	return result
}

// This function does the same as the RepeatingXorKey
// but takes the input and key as a slice of bytes instead of string
func RepeatingXorBytes(input, key []byte) []byte {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = input[i] ^ key[i%len(key)]
	}
	return result
}

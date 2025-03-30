package setOne

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

func RepeatingXorBytes(input, key []byte) []byte {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = input[i] ^ key[i%len(key)]
	}
	return result
}

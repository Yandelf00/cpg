package setOne

import (
	"crypto/aes" // Provides AES cipher functionality
	"fmt"
	"slices"
)

// AesEcbDecrypt decrypts AES-encrypted data in ECB mode using a given key.
// It returns the decrypted bytes or an error if cipher initialization fails.
func AesEcbDecrypt(inpt, key []byte) ([]byte, error) {
	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err // Return error if key size is invalid
	}

	// Prepare a slice to store decrypted output (same length as input)
	rest := make([]byte, len(inpt))

	// Decrypt each 16-byte block (AES block size is 16 bytes)
	for i, j := 0, 16; i < len(inpt); i, j = i+16, j+16 {
		// Decrypt the block from input[i:j] into rest[i:j]
		block.Decrypt(rest[i:j], inpt[i:j])
	}

	return rest, nil
}

// DetectAes tries to detect AES encryption in ECB mode in a given byte slice.
// It does so by checking for repeated 16-byte blocks in the input.
func DetectAes(inpt []byte, lnm int) {
	line := 0 // Holds the line number where ECB mode is detected

	// Iterate over all 16-byte blocks in the input
	for i, j := 0, 16; i < len(inpt); i, j = i+16, j+16 {
		// Compare current block with all subsequent blocks
		for k, l := i+16, j+16; k < len(inpt); k, l = k+16, l+16 {
			if slices.Equal(inpt[i:j], inpt[k:l]) {
				// If two identical blocks are found, assume ECB mode
				line = lnm
			}
		}
	}

	// Output the line number if ECB mode is detected
	fmt.Println(line)
}

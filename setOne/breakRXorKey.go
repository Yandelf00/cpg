package setOne

import (
	"fmt"
	"math/bits"
)

// BreakRKXor attempts to break a repeating-key XOR cipher.
// It finds the 3 most likely key sizes based on normalized Hamming distance,
// then transposes the ciphertext blocks and solves each column as a single-byte XOR.
func BreakRKXor(inpt []byte) {
	// Stores the 3 lowest normalized Hamming distances
	normalizedDistances := []int{100, 100, 100}
	// Corresponding key sizes for the above distances
	results := []int{0, 0, 0}
	// Stores recovered keys for the 3 key sizes
	keys := [][]byte{}
	maxKeysize := 40

	// Loop over candidate key sizes (from 2 to 40)
	for k := 2; k < maxKeysize && 4*k <= len(inpt); k++ {
		totalDist := 0
		numPairs := 0

		// Create offsets to generate multiple block pairs
		offsets := []int{0, k, 2 * k, 3 * k, 4 * k, 5 * k, 6 * k, 7 * k, 8 * k}
		for _, offset := range offsets {
			if offset+2*k > len(inpt) {
				continue
			}
			firstBlock := inpt[offset : offset+k]
			secondBlock := inpt[offset+k : offset+2*k]
			totalDist += HammingDistance(firstBlock, secondBlock)
			numPairs++
		}

		// Normalize the total Hamming distance by (numPairs * key size)
		resNormalised := totalDist / (numPairs * k)

		// Insert the result into the sorted list of top 3 key sizes
		for i := 0; i < 3; i++ {
			if resNormalised < normalizedDistances[i] {
				for j := 2; j > i; j-- {
					normalizedDistances[j] = normalizedDistances[j-1]
					results[j] = results[j-1]
				}
				normalizedDistances[i] = resNormalised
				results[i] = k
				break
			}
		}
	}

	// For each of the top 3 candidate key sizes
	for _, el := range results {
		fmt.Println(el)
		// Break ciphertext into blocks and transpose
		t_result := ProcessBlocks(inpt, el)
		key_tmp := []byte{}
		// Solve each column as a single-byte XOR
		for _, r_el := range t_result {
			_, _, key := BhattacharyyaSingleByteXORCipherByte(r_el)
			key_tmp = append(key_tmp, key)
		}
		keys = append(keys, key_tmp)
	}

	// Decrypt and print results for each found key
	for _, key := range keys {
		decrypted := RepeatingXorBytes(inpt, key)
		fmt.Printf("Decrypted: %s\n", decrypted)
		fmt.Println("the key :", string(key))
	}
}

// ProcessBlocks splits the input into blocks of size `blockSize`,
// and then transposes them to prepare for per-column analysis.
func ProcessBlocks(inpt []byte, blockSize int) [][]byte {
	var result [][]byte
	fullSize := 0
	// Split the input into blocks of `blockSize`
	for k := 0; k+blockSize <= len(inpt); k += blockSize {
		result = append(result, inpt[k:k+blockSize])
		fullSize++
	}
	// Transpose the matrix so each slice contains bytes at the same key position
	t_result := Transpose(result, blockSize, fullSize)

	return t_result
}

// Transpose flips a 2D byte matrix (blockSize x fullSize) into (fullSize x blockSize)
func Transpose(mtrx [][]byte, blockSize int, fullSize int) [][]byte {
	result := [][]byte{}

	// Collect each column of the matrix as a new row
	for j := 0; j < blockSize; j++ {
		tmp := []byte{}
		for i := 0; i < fullSize; i++ {
			tmp = append(tmp, mtrx[i][j])
		}
		result = append(result, tmp)
	}

	return result
}

// HammingDistance calculates the number of differing bits between two byte slices
func HammingDistance(inptOne, inptTwo []byte) int {
	n := len(inptOne)
	res := 0

	// XOR each byte and count the number of set bits (differences)
	for i := 0; i < n; i++ {
		tst := bits.OnesCount8(inptOne[i] ^ inptTwo[i])
		res += tst
	}
	return res
}

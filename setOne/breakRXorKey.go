package setOne

import (
	"fmt"
	"math/bits"
)

func BreakRKXor(inpt []byte) {
	normalizedDistances := []int{100, 100, 100}
	results := []int{0, 0, 0}
	keys := [][]byte{}
	maxKeysize := 40

	for k := 2; k < maxKeysize && 4*k <= len(inpt); k++ {
		totalDist := 0
		numPairs := 0
		offsets := []int{0, k, 2 * k, 3 * k, 4 * k, 5 * k, 6 * k, 7 * k}
		for _, offset := range offsets {
			if offset+2*k > len(inpt) {
				continue
			}
			firstBlock := inpt[offset : offset+k]
			secondBlock := inpt[offset+k : offset+2*k]
			totalDist += HammingDistance(firstBlock, secondBlock)
			numPairs++
		}
		// fmt.Println("key totaldist", numPairs, k, totalDist)
		// resNormalisedfloat := float64(float64(totalDist) / (float64(numPairs) * float64(k)))
		resNormalised := totalDist / (numPairs * k)
		// fmt.Printf("resNormalised %f\n", resNormalisedfloat)
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

	for _, el := range results {
		fmt.Println(el)
		// t_result := ProcessBlocks(inpt, el)
		// key_tmp := []byte{}
		// for _, r_el := range t_result {
		// 	_, _, key := BhattacharyyaSingleByteXORCipherByte(r_el)
		// 	key_tmp = append(key_tmp, key)
		// }
		// keys = append(keys, key_tmp)
	}
	t_result := ProcessBlocks(inpt, results[2])
	key_tmp := []byte{}
	for _, r_el := range t_result {
		_, _, key := BhattacharyyaSingleByteXORCipherByte(r_el)
		key_tmp = append(key_tmp, key)
	}
	keys = append(keys, key_tmp)

	for _, key := range keys {
		for _, el := range key {
			fmt.Println("id", el)
		}
		decrypted := RepeatingXorBytes(inpt, key)
		fmt.Printf("Decrypted: %s\n", decrypted)
		fmt.Println("the key :", string(key))
	}
	// decrypted := RepeatingXorBytes(inpt, keys[0])
	// fmt.Printf("Decrypted: %s\n", decrypted)
}

func ProcessBlocks(inpt []byte, blockSize int) [][]byte {
	var result [][]byte
	fullSize := 0
	for k := 0; k+blockSize <= len(inpt); k += blockSize {
		result = append(result, inpt[k:k+blockSize])
		fullSize++
	}
	t_result := Transpose(result, blockSize, fullSize)

	return t_result
}

func Transpose(mtrx [][]byte, blockSize int, fullSize int) [][]byte {
	result := [][]byte{}

	for j := 0; j < blockSize; j++ {
		tmp := []byte{}
		for i := 0; i < fullSize; i++ {
			tmp = append(tmp, mtrx[i][j])
		}
		result = append(result, tmp)
	}

	return result
}

func HammingDistance(inptOne, inptTwo []byte) int {
	// dstOne := []byte(inptOne)
	// dstTwo := []byte(inptTwo)
	n := len(inptOne)
	res := 0

	for i := 0; i < n; i++ {
		tst := bits.OnesCount8(inptOne[i] ^ inptTwo[i])
		res += tst
	}
	return res
}

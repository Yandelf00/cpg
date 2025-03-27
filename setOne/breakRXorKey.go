package setOne

import (
	"math/bits"
)

func BreakRKXor(inpt []byte) {
	normalizedDistances := []int{100, 100, 100}
	results := []int{0, 0, 0}
	maxKeysize := 60

	for k := 2; k < maxKeysize && 4*k <= len(inpt); k++ {
		totalDist := 0
		numPairs := 0
		offsets := []int{0, k, 2 * k, 3 * k}
		for _, offset := range offsets {
			if offset+2*k > len(inpt) {
				continue
			}
			firstBlock := inpt[offset : offset+k]
			secondBlock := inpt[offset+k : offset+2*k]
			totalDist += HammingDistance(firstBlock, secondBlock)
			numPairs++
		}

		resNormalised := totalDist / (numPairs * k)

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

}

func Transpose(inpt []byte, blockSize int) {

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

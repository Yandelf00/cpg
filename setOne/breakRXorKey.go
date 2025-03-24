package setOne

import (
	"fmt"
	"math/bits"
)

func HammingDistance(inptOne, inptTwo string) {
	dstOne := []byte(inptOne)
	dstTwo := []byte(inptTwo)
	n := len(dstOne)
	res := 0

	for i := 0; i < n; i++ {
		tst := bits.OnesCount8(dstOne[i] ^ dstTwo[i])
		res += tst
	}
	fmt.Println(res)
}

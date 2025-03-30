package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yandelf00/cpg/setOne"
)

// func solveBc() {
// 	var result string
// 	var resultScore float64
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		decrypted, decryptedScore := setOne.BhattacharyyaSingleByteXORCipher(scanner.Text())
// 		if decryptedScore > resultScore {
// 			resultScore = decryptedScore
// 			result = decrypted
// 		}
// 	}
// 	fmt.Println(result)

// 	fmt.Println("Decryption complete. Results written to output.txt.")
// }

// var burning = `Burning 'em, if you ain't quick and nimble
// I go crazy when I hear a cymbal`

// func solveRepeatingXorKey() {
// 	var test []byte
// 	test = setOne.RepeatingXorKey(burning, "ICE")

// 	for _, el := range test {
// 		fmt.Printf("%02x", el)
// 	}
// }

func main() {
	data, err := os.ReadFile("breakRKXORinpt.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataDecoded, err := setOne.Base64Decode(data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", dataDecoded) // Prints raw hex (e.g., "08584826...")

	// setOne.BreakRKXor(dataDecoded)

	// inptOne := "this is a test"
	// inptTwo := "wokka wokka!!!"
	// setOne.HammingDistance(inptOne, inptTwo)
}

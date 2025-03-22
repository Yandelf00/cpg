package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Yandelf00/cpg/setOne"
)

func main() {
	var result string
	var resultScore float64
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		decrypted, decryptedScore := setOne.BhattacharyyaSingleByteXORCipher(scanner.Text())
		if decryptedScore > resultScore {
			resultScore = decryptedScore
			result = decrypted
		}
	}
	fmt.Println(result)

	fmt.Println("Decryption complete. Results written to output.txt.")
}

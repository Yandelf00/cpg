package main

import "github.com/Yandelf00/cpg/setOne"

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
	inptOne := "this is a test"
	inptTwo := "wokka wokka!!!"
	setOne.HammingDistance(inptOne, inptTwo)
}

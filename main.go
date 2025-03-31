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
// 		decrypted, decryptedScore, _ := setOne.BhattacharyyaSingleByteXORCipher(scanner.Text())
// 		if decryptedScore > resultScore {
// 			resultScore = decryptedScore
// 			result = decrypted
// 		}
// 	}
// 	fmt.Println(result)

// 	fmt.Println("Decryption complete. Results written to output.txt.")
// }

// func solveBcByte() {
// 	var result []byte
// 	var resultScore float64
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		pompompom, err := setOne.HexDecode(scanner.Text())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		decrypted, decryptedScore, _ := setOne.BhattacharyyaSingleByteXORCipherByte(pompompom)
// 		if decryptedScore > resultScore {
// 			resultScore = decryptedScore
// 			result = decrypted
// 		}
// 	}
// 	fmt.Println(string(result))

// 	fmt.Println("Decryption complete. Results written to output.txt.")
// }

// func solveRepeatingXorKey() []byte {
// 	var test []byte
// 	test = setOne.RepeatingXorKey(burning, "ICE")

// 	for _, el := range test {
// 		fmt.Printf("%02x", el)
// 	}

// 	return test
// }

// var burning = `Burning 'em, if you ain't quick and nimble
// I go crazy when I hear a cymbal`

// var bcsolvetestext = "this is test and i need it to work"
// var bcsolvetestextEncoded = `=! :i :i=,:=i('-i i',,-i =i=&i>&;"` //encoded with I

// var testtwo = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func main() {
	// pompom, err := setOne.HexDecode(testtwo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// decrypted, _, _ := setOne.BhattacharyyaSingleByteXORCipherByte(pompom)

	// decryptedtwo, _, _ := setOne.BhattacharyyaSingleByteXORCipher(testtwo)
	// fmt.Println("decrypted two ", decryptedtwo)
	// fmt.Println("decrypted", string(decrypted))
	// solveBc()
	// solveBcByte()

	// fmt.Println("first loop")
	// for i := 0; i <= 20; i++ {
	// 	res, err := setOne.FixedXorKey(inputtxt, byte(i))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(res)

	// }

	// fmt.Println("second loop")
	// for i := 0; i <= 20; i++ {
	// 	res, err := setOne.FixedXorKeyByte(pompom, byte(i))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(res)
	// }

	// key := 'Y'
	// res, err := setOne.FixedXorKeyByte([]byte(burning), byte(key))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(res))
	// resDec, err := setOne.FixedXorKeyByte(res, byte(key))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(resDec))
	data, err := os.ReadFile("breakRKXORinpt.txt")
	// data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	dataDecoded, err := setOne.Base64Decode(data)
	fmt.Println(dataDecoded)

	if err != nil {
		log.Fatal(err)
	}

	setOne.BreakRKXor(string(dataDecoded))
	// for i := 0; i <= 255; i++ {
	// 	fmt.Println(byte(i))
	// }
	// inptOne := "this is a test"
	// inptTwo := "wokka wokka!!!"
	// setOne.HammingDistance(inptOne, inptTwo)
}

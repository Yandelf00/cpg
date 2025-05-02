package main

import (
	"fmt"

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

	//testing single byte xor cipher 
	res, _, _:= setOne.BhattacharyyaSingleByteXORCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	fmt.Println(res)

	//testing FixedXor
	// test, err := setOne.FixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(test)
	// fmt.Println(string(test))
	// fmt.Printf("%x", test)
	//detect aes in ecb mode
	// lnm := 0
	// file, err := os.Open("detectaes.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	//
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	lnm++
	// 	aesDecoded, err := setOne.HexDecode(scanner.Text())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	setOne.DetectAes(aesDecoded, lnm)
	// }
	//
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	//decoding aes in acb mode
	// dataAes, err := os.ReadFile("aesecbdecrypt.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// dataAesDecoded, err := setOne.Base64Decode([]byte(dataAes))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// key := []byte("YELLOW SUBMARINE")
	// res, err := setOne.AesEcbDecrypt(dataAesDecoded, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(res))
	//------------------------------------------------

	// test := []byte{29, 77, 2, 19, 26, 31, 14, 1, 78, 22, 0, 73, 84, 67, 78}
	// fmt.Println(string(test))

	// decrypted, _, _ := setOne.BhattacharyyaSingleByteXORCipherByte(test)
	// fmt.Println(string(decrypted))
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
	// data, err := os.ReadFile("breakRKXORinpt.txt")
	// // data, err := os.ReadFile("test.txt")
	// // data, err := os.ReadFile("testtwo.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// inpt := "Repeating XOR is tricky but solvable with the right approach!"

	// xoredinpt := setOne.RepeatingXorBytes([]byte(inpt), []byte("GOLANG"))
	// fmt.Println(string(xoredinpt))

	// dataEncoded := setOne.Base64Encode([]byte(xoredinpt))

	// fmt.Println(string(dataEncoded))
	// data := "Kg4XFFQbBBFfQh1bQh0XU0IcF1QeCxVfQh1bQh0XU0IcF1QeCxVfQh1bQh0XU0IcF1QeCxVfQh1bQh0XU0IcF1QeCxVfQh1bQh0XU0IcF1QeCxVfQh1bQh0XU0IcF1QeCw=="

	// dataDecoded, err := setOne.Base64Decode([]byte(data))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := setOne.RepeatingXorBytes(dataDecoded, []byte("GOLANG"))
	// fmt.Println(string(result))

	// setOne.BreakRKXor(dataDecoded)
	// for i := 0; i <= 255; i++ {
	// 	fmt.Println(byte(i))
	// }
	// inptOne := "this is a test"
	// inptTwo := "wokka wokka!!!"
	// setOne.HammingDistance(inptOne, inptTwo)
}

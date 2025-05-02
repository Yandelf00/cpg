package setOne

import (
	"log"
	"math"
	"strings"
)

//Struct that holds the decrypted phrase,
//its Bhattacharyya coefficient(score),
//and the key with which it was decrypted.
type BhattaRes struct {
	phrase string
	score  float64
	key    byte
}


//This function uses bhattacharyya coefficient
//to find the decrypted message, it takes a string input.
//The point of this function is to detect the english sentence
//because we don't wanna look at all the possible outputs and 
//find the english sentence by ourselves
func BhattacharyyaSingleByteXORCipher(input string) (string, float64, byte) {
	var bc float64
	bhattaRes := BhattaRes{
		phrase: "",
		score:  0.0,
		key:    0,
	} //initialization of the result

	for i := 0; i <= 255; i++ {
		res, err := FixedXorKey(input, byte(i)) //xor our input against every possible character
		if err != nil {
			log.Fatal(err)
		}
		bc = BhattaPhraseAnalysis(string(res)) //the function that calculates the bc for each input
		if bc > bhattaRes.score {
			bhattaRes.phrase = string(res)
			bhattaRes.score = bc
			bhattaRes.key = byte(i)
		} //if the bc of the current phrase is higher than the one stored, we take the new one
	}

	return bhattaRes.phrase, bhattaRes.score, bhattaRes.key
}


//This function calculates the frequency of the letters in our decoded string (non normalized)
func BhattaCalculateFrequency(phrase string) map[rune]float64 {
	var frequencies = map[rune]float64{
		'a': 0, 'b': 0, 'c': 0, 'd': 0,
		'e': 0, 'f': 0, 'g': 0, 'h': 0,
		'i': 0, 'j': 0, 'k': 0, 'l': 0,
		'm': 0, 'n': 0, 'o': 0, 'p': 0,
		'q': 0, 'r': 0, 's': 0, 't': 0,
		'u': 0, 'v': 0, 'w': 0, 'x': 0,
		'y': 0, 'z': 0, ' ': 0,
	} //initates the map of frequencies for our input
	totalLetter := 0

	for _, char := range strings.ToLower(phrase) {
		if (char > 'a' && char < 'z') || char == ' ' {
			frequencies[char]++
			totalLetter++
		}
	} //goes through every letter, adds the frequency if the letter is english, and add to total letters

	for char := range frequencies {
		frequencies[char] = (frequencies[char] / float64(totalLetter)) * 100
	}//gives the percentage of each letter
	return frequencies
}


func BhattaPhraseAnalysis(phrase string) float64 {
	frequency := BhattaCalculateFrequency(phrase)
	normalizedFrequency := NormalizeFrequency(frequency)
	normalizedEndlishFreq := NormalizeFrequency(englishFrequencies)

	bc := 0.0

	for char, observed := range normalizedFrequency {
		expected, exists := normalizedEndlishFreq[char]
		if exists {
			bc += math.Sqrt(observed * expected)
		}
	}

	return bc
}

func NormalizeFrequency(input map[rune]float64) map[rune]float64 {
	total := 0.0
	for _, freq := range input {
		total += freq
	}

	normalizedFreq := make(map[rune]float64)
	for char, freq := range input {
		normalizedFreq[char] = freq / total
	}

	return normalizedFreq
}

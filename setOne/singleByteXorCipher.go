package setOne

import (
	"fmt"
	"log"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/abadojack/whatlanggo"
)

type ZscoreRes struct {
	phrase string
	score  float64
}

type BhattaRes struct {
	phrase string
	score  float64
}

var englishFrequencies = map[rune]float64{
	'a': 8.167, 'b': 1.492, 'c': 2.782, 'd': 4.253,
	'e': 12.702, 'f': 2.228, 'g': 2.015, 'h': 6.094,
	'i': 6.966, 'j': 0.153, 'k': 0.772, 'l': 4.025,
	'm': 2.406, 'n': 6.749, 'o': 7.507, 'p': 1.929,
	'q': 0.095, 'r': 5.987, 's': 6.327, 't': 9.056,
	'u': 2.758, 'v': 0.978, 'w': 2.36, 'x': 0.15,
	'y': 1.974, 'z': 0.074, ' ': 19.0,
}

func CalculateInputFrequency(phrase string) map[rune]float64 {
	frequencies := make(map[rune]float64)
	totalLetter := 0

	for _, char := range strings.ToLower(phrase) {
		if (char > 'a' && char < 'z') || char == ' ' {
			frequencies[char]++
			totalLetter++
		}
	}

	for char := range frequencies {
		frequencies[char] = (frequencies[char] / float64(totalLetter)) * 100
	}
	return frequencies
}

func CalculateZScore(observed, expected, stdDev float64) float64 {
	return (observed - expected) / stdDev
}

func PhraseAnalysis(phrase string) map[rune]float64 {
	frequency := CalculateInputFrequency(phrase)
	zScore := make(map[rune]float64)

	stdDev := 1.0

	for char, observed := range frequency {
		expected, exists := englishFrequencies[char]
		if exists {
			zScore[char] = CalculateZScore(observed, expected, stdDev)
			continue
		}
		zScore[char] = 0.0
	}

	return zScore
}

func CalulatePhraseScore(zScore map[rune]float64) float64 {
	var phraseScore float64 = 0

	for _, z := range zScore {
		phraseScore += math.Pow(z, 2)
	}

	return phraseScore
}

func SingleByteXORCipher(input string) string {
	var alpha int = 65
	var decrypted []string
	var decryptedZScore map[rune]float64
	var decryptedPhraseScore float64
	res := ZscoreRes{
		phrase: "",
		score:  1000.0,
	}

	for i := 0; i < 26; i++ {
		res, err := FixedXorKey(input, byte(alpha+i))
		if err != nil {
			log.Fatal(err)
		}
		decrypted = append(decrypted, string(res))
	}

	for _, phrase := range decrypted {
		if !utf8.ValidString(phrase) {
			continue
		}
		decryptedZScore = PhraseAnalysis(phrase)
		decryptedPhraseScore = CalulatePhraseScore(decryptedZScore)
		if decryptedPhraseScore < res.score {
			res.phrase = phrase
			res.score = decryptedPhraseScore
		}
	}
	return res.phrase
}

func SyntheticSingleByteXORCipher(input string) string {
	var alpha int = 65
	var decrypted []string
	info := whatlanggo.Detect("")
	var result string

	for i := 0; i < 26; i++ {
		res, err := FixedXorKey(input, byte(alpha+i))
		if err != nil {
			log.Fatal(err)
		}
		decrypted = append(decrypted, string(res))
	}

	for _, phrase := range decrypted {
		info = whatlanggo.Detect(phrase)
		if info.Lang == whatlanggo.Eng {
			result = phrase
			fmt.Println("this is in english", result)
		}
	}
	return result
}

func BhattacharyyaSingleByteXORCipher(input string) string {
	var decrypted []string
	var bc float64
	bhattaRes := BhattaRes{
		phrase: "",
		score:  0.0,
	}

	for i := 0; i <= 255; i++ {
		res, err := FixedXorKey(input, byte(i))
		if err != nil {
			log.Fatal(err)
		}
		decrypted = append(decrypted, string(res))
	}

	for _, phrase := range decrypted {
		bc = BhattaPhraseAnalysis(phrase)
		if bc > bhattaRes.score {
			bhattaRes.phrase = phrase
			bhattaRes.score = bc
		}
	}

	return bhattaRes.phrase
}

func BhattaCalculateFrequency(phrase string) map[rune]float64 {
	var frequencies = map[rune]float64{
		'a': 0, 'b': 0, 'c': 0, 'd': 0,
		'e': 0, 'f': 0, 'g': 0, 'h': 0,
		'i': 0, 'j': 0, 'k': 0, 'l': 0,
		'm': 0, 'n': 0, 'o': 0, 'p': 0,
		'q': 0, 'r': 0, 's': 0, 't': 0,
		'u': 0, 'v': 0, 'w': 0, 'x': 0,
		'y': 0, 'z': 0, ' ': 0,
	}
	totalLetter := 0

	for _, char := range strings.ToLower(phrase) {
		if (char > 'a' && char < 'z') || char == ' ' {
			frequencies[char]++
			totalLetter++
		}
	}

	for char := range frequencies {
		frequencies[char] = (frequencies[char] / float64(totalLetter)) * 100
	}
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

func BruteSingleByteXORCipher(input string) string {
	validChars := map[rune]struct{}{
		'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {}, 'h': {},
		'i': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {}, 'n': {}, 'o': {}, 'p': {},
		'q': {}, 'r': {}, 's': {}, 't': {}, 'u': {}, 'v': {}, 'w': {}, 'x': {},
		'y': {}, 'z': {}, ' ': {}, '"': {}, '\'': {}, '.': {}, ',': {}, '?': {}, '!': {},
		'\n': {},
	}
	// var alpha int = 65
	var decrypted []string

	for i := 0; i <= 255; i++ {
		res, err := FixedXorKey(input, byte(i))
		if err != nil {
			log.Fatal(err)
		}
		decrypted = append(decrypted, string(res))
	}

	for _, str := range decrypted {
		if isValidString(str, validChars) {
			return str
		}
	}
	return ""
}

func isValidString(str string, validChars map[rune]struct{}) bool {
	for _, char := range strings.ToLower(str) {
		if _, exists := validChars[char]; !exists {
			return false
		}
	}
	return true
}

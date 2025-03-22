package setOne

import (
	"log"
	"math"
	"strings"
	"unicode/utf8"
)

type ZscoreRes struct {
	phrase string
	score  float64
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

func ZScoreSingleByteXORCipher(input string) string {
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

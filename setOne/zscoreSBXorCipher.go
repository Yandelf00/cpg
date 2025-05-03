package setOne

import (
	"log"
	"math"
	"strings"
	"unicode/utf8"
)

// ZscoreRes holds a decrypted phrase and its associated z-score (used for scoring how "English-like" it is).
type ZscoreRes struct {
	phrase string
	score  float64
}

// CalculateZScore computes the standard z-score given observed, expected, and std deviation values.
func CalculateZScore(observed, expected, stdDev float64) float64 {
	return (observed - expected) / stdDev
}

// PhraseAnalysis computes the z-score of each rune (letter or space) in the phrase
// compared to standard English letter frequencies.
func PhraseAnalysis(phrase string) map[rune]float64 {
	frequency := CalculateInputFrequency(phrase) // Get frequency % of each char in input
	zScore := make(map[rune]float64)

	stdDev := 1.0 // Assumes unit standard deviation for simplicity

	for char, observed := range frequency {
		expected, exists := englishFrequencies[char] // Look up expected frequency
		if exists {
			zScore[char] = CalculateZScore(observed, expected, stdDev)
			continue
		}
		zScore[char] = 0.0 // If char isn't in reference table, assign 0
	}

	return zScore
}

// CalulatePhraseScore returns the sum of squared z-scores (lower means closer to English distribution).
func CalulatePhraseScore(zScore map[rune]float64) float64 {
	var phraseScore float64 = 0

	for _, z := range zScore {
		phraseScore += math.Pow(z, 2)
	}

	return phraseScore
}

// ZScoreSingleByteXORCipher tries all A-Z single-byte keys to decrypt the input string,
// then scores each output using z-score analysis to find the most English-like result.
func ZScoreSingleByteXORCipher(input string) string {
	var alpha int = 65     // ASCII value for 'A'
	var decrypted []string // Stores decrypted candidates
	var decryptedZScore map[rune]float64
	var decryptedPhraseScore float64

	// Initialize result with a high score
	res := ZscoreRes{
		phrase: "",
		score:  1000.0,
	}

	// Try XORing with every uppercase letter A-Z (ASCII 65–90)
	for i := 0; i < 26; i++ {
		res, err := FixedXorKey(input, byte(alpha+i))
		if err != nil {
			log.Fatal(err)
		}
		decrypted = append(decrypted, string(res))
	}

	// Analyze and score each decrypted phrase
	for _, phrase := range decrypted {
		if !utf8.ValidString(phrase) {
			continue // Skip invalid UTF-8 strings
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

// CalculateInputFrequency computes the percentage frequency of each valid character
// (a–z and space) in the input phrase.
func CalculateInputFrequency(phrase string) map[rune]float64 {
	frequencies := make(map[rune]float64)
	totalLetter := 0

	// Count letters and spaces only (case-insensitive)
	for _, char := range strings.ToLower(phrase) {
		if (char > 'a' && char < 'z') || char == ' ' {
			frequencies[char]++
			totalLetter++
		}
	}

	// Convert counts to percentages
	for char := range frequencies {
		frequencies[char] = (frequencies[char] / float64(totalLetter)) * 100
	}
	return frequencies
}

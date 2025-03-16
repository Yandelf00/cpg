package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/abadojack/whatlanggo"
)

// isMeaningfulText checks if the input contains enough printable/graphic characters
func isMeaningfulText(input string) bool {
	printableCount := 0
	for _, char := range input {
		if unicode.IsPrint(char) && unicode.IsGraphic(char) {
			printableCount++
		}
	}
	// Require at least 50% of the characters to be printable/graphic
	return float64(printableCount)/float64(utf8.RuneCountInString(input)) >= 0.5
}

func main() {
	input := "T◄§PåqjF�Zö0Ll�Bå↕�♦n3�_w"

	// Preprocess the input to check if it's meaningful
	if !isMeaningfulText(input) {
		fmt.Println("Input contains too many non-printable or non-graphic characters.")
		return
	}

	// Detect the language
	info := whatlanggo.Detect(input)
	if info.Lang == whatlanggo.Eng {
		fmt.Println("This is in English.")
	} else {
		fmt.Println("It ain't in English.")
	}
}

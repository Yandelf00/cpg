package setOne

import (
	"log"
	"strings"
)

var validChars = map[rune]struct{}{
	'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {}, 'h': {},
	'i': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {}, 'n': {}, 'o': {}, 'p': {},
	'q': {}, 'r': {}, 's': {}, 't': {}, 'u': {}, 'v': {}, 'w': {}, 'x': {},
	'y': {}, 'z': {}, ' ': {}, '"': {}, '\'': {}, '.': {}, ',': {}, '?': {}, '!': {},
	'\n': {},
}

func BruteSingleByteXORCipher(input string) string {
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

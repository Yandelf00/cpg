package settwo

import (
	"errors"
)

// PKCSPadding function implements the padding scheme PKCS#7
// this algorithm takes an input and padd it until it reaches
// the desired length, if the difference between the current goal
// and the desired one is lets say 5, its gonna add 5 "\x05" to
// our input etc..
//
// Parameters :
// -inpt : this is the input string that we are going to pad
// -goal : this is the length we want to reach after padding our input
//
// Returns :
// -result : our padded result return in form of bytes (to get the desired
// string output it should be printed using fmt.Printf("%q\n", result))
func PKCSPadding(inpt string, goal int) ([]byte, error) {
	if goal < len(inpt) {
		return nil, errors.New("goal can't be lower than input length")
	}

	result := []byte(inpt)
	paddingLength := goal - len(inpt)
	plByte := byte(paddingLength)

	//adds as many bytes as needed to reach goal
	for i := 0; i < paddingLength; i++ {
		result = append(result, plByte)
	}

	return result, nil
}

package setOne

// This is a global variable that contains the
// english letter frequencies
// this variable is wrong tho, as the sum of
// frequencies is 118, thats why i normalize it
// in another function, and it's actually better to
// use correct frequencies and have the normalization ready
// so it's more efficient(but it works with my algorithms so
// i am keeping it)
var englishFrequencies = map[rune]float64{
	'a': 8.167, 'b': 1.492, 'c': 2.782, 'd': 4.253,
	'e': 12.702, 'f': 2.228, 'g': 2.015, 'h': 6.094,
	'i': 6.966, 'j': 0.153, 'k': 0.772, 'l': 4.025,
	'm': 2.406, 'n': 6.749, 'o': 7.507, 'p': 1.929,
	'q': 0.095, 'r': 5.987, 's': 6.327, 't': 9.056,
	'u': 2.758, 'v': 0.978, 'w': 2.36, 'x': 0.15,
	'y': 1.974, 'z': 0.074, ' ': 19.0,
}

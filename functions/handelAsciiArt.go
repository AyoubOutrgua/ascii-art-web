package functions

func HandelAsciiArt(inputText, banner string) (string, bool) {
	// check range of printable characters
	if !CheckRange(inputText) {
		return "", true
	}

	// check the validity of the banner
	if !CheckBanner(banner) {
		return "", true
	}

	// split input text to slice of string
	wordsSlice := SplitInput(inputText)

	// calling functions to handle the input
	return AppendArt(wordsSlice, AsciiArtTable(banner)), false
}

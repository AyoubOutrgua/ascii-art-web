package functions

import (
	"fmt"
)

func HandelAsciiArt(inputText, banner string) string {
	// check range of printable characters
	if !CheckRange(inputText) {
		fmt.Println("error, string is not valid ")
		return ""
	}

	// check the validity of the banner
	if !CheckBanner(banner) {
		fmt.Println(banner)
		fmt.Println("error, banner is not valid ")
		return ""
	}

	// split input text to slice of string
	wordsSlice := SplitInput(inputText)

	// calling functions to handle the input
	return AppendArt(wordsSlice, AsciiArtTable(banner))
}

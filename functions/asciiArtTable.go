package functions

import (
	"fmt"
	"os"
	"strings"
)

// a function to read the file and store the ascii values in a map
func AsciiArtTable(banner string) map[int]string {
	bannerName := "banners/" + banner + ".txt"
	fileData, err := os.ReadFile(bannerName)
	if err != nil {
		fmt.Println("error:", err)
		return map[int]string{}
	}
	blocks := strings.Split(string(fileData), "\n\n")
	// after split formatting if there is an extra new line
	if strings.Count(blocks[0], "\n") == 8 {
		blocks[0] = blocks[0][1:]
	}
	// initializing a map and storing the ascii values in
	asciiArtTable := make(map[int]string)
	key := 32
	for _, block := range blocks {
		asciiArtTable[key] = block
		key++
	}
	return asciiArtTable
}

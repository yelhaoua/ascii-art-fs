package asciiart

import (
	"fmt"
	"strings"
)

func output(str string, array [][]string) string {
	out := ""
	// start looping to8 the hieght of line
	for i := 0; i < 8; i++ {
		for _, fin := range str {
			// skip the unprintibal chachacters
			if fin > 126 || fin < 32 {
				continue
			}
			out += array[int(rune(fin)-32)][i]
		}
		if i != 7 {
			out += "\n"
		}

	}
	return out
}

func PrintSymbole(array [][]string, word string) string {
	
	out := ""
	// check if the word not empty and the array of banner
	if word == "" || len(array) == 0 {
		return ""
	}
	// change the \n as string to newline
	str := strings.ReplaceAll(word, `\n`, "\n")
	fmt.Println(str)

	// skip addional newline in print
	if strings.Trim(str, "\n") == "" {

		return ""
	}
	// spliting the word
	words := strings.Split(word, "\n")

	// Print each word
	for _, s := range words {
		out = output(s, array)
	}
	return out
}

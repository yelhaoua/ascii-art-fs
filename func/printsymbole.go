package asciiart

import (
	"strings"
)

func output(str string, array [][]string) string {
	text := ""
	// start looping to8 the hieght of line
	for i := 0; i < 8; i++ {
		for _, fin := range str {
			// skip the unprintibal chachacters
			if fin > 126 || fin < 32 {
				continue
			}
			text += array[int(rune(fin)-32)][i]
		}
		if i != 7 {

			text += "\n"
		}

	}
	return text
}

func PrintSymbole(array [][]string, word string) string {
	res := ""
	// check if the word not empty and the array of banner
	if word == "" || len(array) == 0 {
		return ""
	}
	// change the \n as string to newline
	str := strings.ReplaceAll(word, `\n`, "\n")

	// skip addional newline in print
	if strings.Trim(str, "\n") == "" {

		return ""
	}
	// spliting the word
	words := strings.Split(word, "\n")

	// Print each word
	for _, s := range words {
		res = output(s, array)
	}
	return res
}

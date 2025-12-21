package asciiart

import (
	"strings"
)

func processLine(line string, banner [][]string) string {
	if line == "" {
		return "\n"
	}

	var out string
	// start looping to8 the hieght of line
	for i := 0; i < 8; i++ {
		for _, fin := range line {
			out += banner[int(rune(fin)-' ')][i]
		}
		out += "\n"
	}
	return out
}

func PrintSymbole(banner [][]string, input string) string {
	var out string
	// check if the word not empty and the array of banner
	if input == "" || len(banner) == 0 {
		return ""
	}

	// check unprintibal chachacters
	for _, c := range input {
		if c > '~' || c < ' ' {
			return ""
		}
	}

	// spliting the lines
	lines := strings.Split(input, "\\n")
	if lines[0] == "" && lines[1] == "" {
		lines = lines[1:]
	}

	// Print each line
	for _, s := range lines {
		out += processLine(s, banner)
	}
	return out
}

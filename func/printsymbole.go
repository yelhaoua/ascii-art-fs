package asciiart

import (
	"fmt"
	"strings"
)

func PrintSymbole(arr [][]string, woord string) {
	if woord == "" {
		return
	}

	str := strings.ReplaceAll(woord, `\n`, "\n")

	if strings.Trim(str, "\n") == "" {

		for i := 0; i < len(woord); i++ {
			if woord[i] == '\\' {
				fmt.Println()
			}
		}
		return
	}

	words := strings.Split(str, "\n")
	for _, val := range words {
		if val == "" {
			fmt.Println()
			continue
		} else {
			for i := 0; i < 8; i++ {
				for _, sVal := range val {
					if sVal < ' ' || sVal > '~' {
						continue
					} else if sVal >= 0 && int(rune(sVal)-32) < len(arr) {
						fmt.Print(arr[int(rune(sVal)-32)][i])
					}
				}
				fmt.Println()
			}
		}
	}
}

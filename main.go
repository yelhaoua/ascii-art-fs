package main

import (
	"fmt"
	"os"

	asciiart "asciiart/func"
)

func main() {
	if len(os.Args) == 2 {
		// here is default banner
		Content := asciiart.Splite("standard")
		// os.args[1] == the user text
		results := asciiart.PrintSymbole(Content, os.Args[1])
		fmt.Println(results)
	} else if len(os.Args) == 3 {
		// os.args[2] == the name of  banner
		Content := asciiart.Splite(os.Args[2])
		// os.args[1] == the user text
		results := asciiart.PrintSymbole(Content, os.Args[1])
		if len(results) != 0 {
			fmt.Println(results)
		}
	} else {
		printErr()
		return
	}
}

func printErr() {
	fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
}

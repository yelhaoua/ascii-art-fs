package main

import (
	"fmt"
	"os"

	asciiart "asciiart/func"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		printErr()
		return
	}

	banner := "standard"
	if len(os.Args) == 3 {
		// os.args[2] == name of banner
		banner = os.Args[2]
	}
	Content := asciiart.Splite(banner)

	// os.args[1] == the user text
	results := asciiart.PrintSymbole(Content, os.Args[1])
	if len(results) != 0 {
		fmt.Print(results)
	} else {
		fmt.Println("Invalid input")
	}
}

func printErr() {
	fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
}

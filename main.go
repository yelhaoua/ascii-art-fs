package main

import (
	asciiart "asciiart/func"
	"os"
)

func main() {
	
	if len(os.Args) == 2 {
		Content := asciiart.Splite("standard")
		// os.args[1] == the user text
		asciiart.PrintSymbole(Content, os.Args[1])
	} else if len(os.Args) == 3 {
		// os.args[2] == te name of  banner
		Content := asciiart.Splite(os.Args[2])
		// os.args[1] == the user text
		asciiart.PrintSymbole(Content, os.Args[1])
	} else {
		return
	}

}

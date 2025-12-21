package main

import (
	asciiart "asciiart/func"
	"fmt"
	"testing"
)

func TestHelloName(t *testing.T) {
	testname := []string{"hello", "Hello There!", "Hello There!"}
	testbanner := []string{"standard", "shadow", "thinkertoy"}
	var res []string
	var arr [][]string
	expected := []string{` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               `,
		`                                                                                         
_|    _|          _| _|                _|_|_|_|_| _|                                  _| 
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| 
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| 
                                                                                         
                                                                                         `,
		`                                                
o  o     o o           o-O-o o                o 
|  |     | |             |   |                | 
O--O o-o | | o-o         |   O--o o-o o-o o-o o 
|  | |-' | | | |         |   |  | |-' |   |-'   
o  o o-o o o o-o         o   o  o o-o o   o-o O 
                                                
                                                `}
									
	for i, v := range testname {
		arr = asciiart.Splite(testbanner[i])
		res = append(res, asciiart.PrintSymbole(arr, v))
	}
	for i := 0; i < len(expected); i++ {
		if res[i] == expected[i] {
			println("Test Pass ")

		} else {

			fmt.Println(res[i])
			println()
			fmt.Println(expected[i])

		}

	}
}

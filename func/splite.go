package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func Splite(fileName string) [][]string {
	var res [][]string
	var all []string
	// Define the path of the banners
	data, err := os.ReadFile("./banners/" + fileName + ".txt")
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}

	// Split the file by \n
	all = strings.Split(string(data), "\n")

	// Make every array contain 8 characters
	for i := 1; i+8 < len(all); i += 9 {
		res = append(res, all[i:i+8])
	}

	fmt.Println(res)
	return res
}

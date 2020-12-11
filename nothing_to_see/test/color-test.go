package main

import (
	"flag"
	"fmt"
)

func main() {
	color:= flag.String("color", "", "put your color in rgb, separate the 3 varaibles with a comma and no spaces")
	flag.Parse()
	fmt.Println(*color)
	var r,g,b string
	if *color != "" {
		var table []string
		var numb =""
		for _, char := range *color {
			if char == ',' {
				table = append(table, numb)
				numb=""
			}else {
				numb+=string(char)
			}
		}
		table = append(table, numb)
		r= table[0]
		b= table [1]
		g= table [2]
	}

	fmt.Printf("\x1b[38;2;"+r+";"+g+";"+b+"m%s\x1b[0m\n", "yo")
}

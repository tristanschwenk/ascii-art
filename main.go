package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args)==1 {
		//If user doesn't put arguments, the script launch the help function and close itself
		help()
		os.Exit(0)
	}else {
		//Initiate the script and keep on a table the lines of the first character of each letters
		startLine := Initialisation(os.Args[len(os.Args)-1])
		//Begin the script
		PrintLine(startLine)
	}
}

func Initialisation(word string) []rune {
	//Get the word and find the first line of all letters of the say word
	start_line := []rune{}
	for _,char := range word {
		val:= 11+9*(char-33)
		start_line=append(start_line, val)
	}
	return start_line
}

func FileRead(fileName string) []byte {
	filename:= "./font/"+fileName
	//Open the right file to use as print base
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	stat, err_stat := file.Stat()
	if err_stat != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	data:=make([]byte, stat.Size())
	_, err_read := file.Read(data)
	if err_read != nil {
		fmt.Println(err.Error())
	}
	file.Close()
	return data
}

func GetColor(color string) []string{
	//Read the argument send with the color flag
	var colors []string
	var numb =""
	//Go threw the value of the color flag, add every character to a string, if it's a comma, append the string to a table
	for _, char := range color {
		if char == ',' {
			colors = append(colors, numb)
			numb=""
		}else {
			numb+=string(char)
		}
	}
	colors = append(colors, numb)
	//Return the table we created to the printing function
	return colors
}

func CheckFlags(output, color string) {
	//Check if those 2 flags are used, if so, we quit the program as we can't print color in a file
	if output !="" && color !="" {
		fmt.Println("Those flags can't be used together")
		os.Exit(1)
	}
}

func PrintLine(startLine []rune)  {
	//Set up the flags
	fs := flag.String("fs", "standard", "You can choose between 3 differents types of font: standart, shadow and thinkertoy")
	outpout:= flag.String("output", "", "Type the outpout flag if you want to print your text in an other file, just add the name of the file and the program will create it for you")
	color:= flag.String("color", "", "Type the color in RGB format, here is an example : -color XXX,XXX,XXX")

	flag.Parse()

	/*fmt.Println("fs:", *fs)
	fmt.Println("output:", *outpout)
	fmt.Println("color:", *color)
	fmt.Println("tail:", flag.Args())*/

	CheckFlags(*outpout, *color)

	//Read the file where all the ASCII characters are stored
	data:= FileRead(*fs+".txt")

	//Create, if needed, a file and declare a variable that will store the content we will create before sending it (see at the end)
	var d1 []byte
	if *outpout != "" {
		f, err_create := os.Create(*outpout)
		if err_create != nil {
			fmt.Println(err_create.Error())
		}
		defer f.Close()
	}

	//Create, if needed, variables holding the color needed
	var colors []string
	if *color != "" {
		colors=GetColor(*color)
	}

	//For each line of a character
	for i := 0; i < 8; i++ {
		for j := 0; j < len(startLine); j++ {
			//Print multiples letter in one line
			var count int
			var temp int

			for k := 0; k < len(data); k++ {
				//We go threw the file, and count the return line (\n)
				if data[k]==10 {
					count++
					if count == int(startLine[j])+i-1 {
						//If our counter is right before the letter we want to print, we implement our temp variable
						temp = count
					}
				}
				if temp != 0 {
					if count == temp+1 {
						//This means that the file read went to a new line, we then have to stop printing
						break
					}else if data[k]!=10{
						//When we are at the right line, we print everything except return line
							//If we've used the output flag, then we append the character to a table
						if *outpout != "" {
							d1 = append(d1, data[k])
							//If we've used the color flag, then we print it with the right color
						}else if *color != "" {
							fmt.Printf("\x1b[38;2;"+colors[0]+";"+colors[1]+";"+colors[2]+"m%s\x1b[0m", string(data[k]))
							// Else we simply print the character
						}else {
							fmt.Printf(string(data[k]))
						}
					}
				}
			}
		}
		//If we've used the output flag, we append a return line to our table
		if *outpout != "" {
			d1=append(d1, byte('\n'))
		// Else we just print a new line
		}else {
			fmt.Println("")
		}
	}
	// Finally, if we've used the output flag, we write our byte table on the file send as parameter of the flag that we created earlier
	if *outpout != "" {
		err := ioutil.WriteFile(*outpout, d1, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func help() {
	welcome:=Initialisation("ASCII-ART")
	PrintLine(welcome)
	fmt.Println("You are on my version of the ascii-art")
	fmt.Println("Keep in mind that the word you want to print has to be the last argument that you put on your command")
	fmt.Println("")
	fmt.Println("It has multiple dependences that you can use:")
	fmt.Println("	-fs flag: You can choose between 3 differents types of font: standart, shadow and thinkertoy")
	fmt.Println("	-outpout flag: Type the outpout flag if you want to print your text in an other file, just add the name of the file and the program will create it for you")
	fmt.Println("	-color flag : Type the color in RGB format, here is an example : -color XXX,XXX,XXX")
}

//func PrintLetter(data []byte, letter int)  {
//	//Print one letter, we put as parameter the line of the beginning of the letter
//	var count int
//	var temp int
//	for i := 0; i < len(data); i++ {
//		if data[i]==10 {
//			count++
//		}
//		if count == letter-1 {
//			temp=count
//		}
//		if temp != 0 {
//			z01.PrintRune(rune(data[i]))
//			if count == temp+9 {
//				break
//			}
//		}
//	}
//}

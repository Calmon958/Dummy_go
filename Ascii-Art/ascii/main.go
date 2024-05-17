package main

// package ascii

import (
	"fmt"
	"os"
)

func main() {
	l := len(os.Args)
	if l != 2 {
		fmt.Println("Enter message to be printed")
		return
	}
	//  else {
	// var file = flag
	// 	// var file string

	// 	file := os.Args[3]
	// 	// os.Args[len(os.Args)] = file
	// 	var output string

	// 	if file == "standard" {
	// 		output = "standard.txt"
	// 	} else if file == "shadow" {
	// 		output = "shadow.txt"
	// 	} else if file == "thinkertoy" {
	// 		output = "thinkertoy.txt"
	// 	} else {
	// 		fmt.Println("File not found. Choose another file")
	// 		os.Exit(1)
	// 	}
	// 	fmt.Print(output)
	// }
	// }

	s, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("text file not found")
		return
	}

	m := asciiArt(string(s))
	res := Tab(os.Args[1])

	Paragraph(res, m)
}

/* to do

account for the possibility that the file doesn't contain all lines(A)
change the type of file that contains the ascii text(W)
account for /v, /a, /f, /b, /r(W)
testfiles for each case(B)
README file update(A)

*/

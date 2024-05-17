package main

import (
	"fmt"
	"strings"
)

// maps the string
func asciiArt(s string) map[rune][]string {
	wordArt := strings.Split(string(s), "\n")
	char := ' ' // declaring first character
	m := make(map[rune][]string)
	var buffer []string
	len := len(wordArt)
	// for storing the ascii the mapped key

	for a := 0; a < len; a++ {
		line := wordArt[a]
		if a == 0 {
			buffer = make([]string, 0)
			/*buffer acts as a holder, has nothing in it
			it stores what is stored in each specific line.
			In this case it's an empty line so nothing*/
		} else if a%9 == 0 {
			m[char] = buffer
			/* stores the character it has mapped from the eight lines to the buffer and looks for its corresponding
			key */
			buffer = make([]string, 0)
			fmt.Println(m)
			// resets the buffer back to nil
			char++
			// after representing the first character it moves to the next one in the provided textfile
		} else {
			buffer = append(buffer, line)
			// stores what is found in each line which isn't nil then moves to the next one
		}
	}
	return m
}


// runeart := [rune][]string
// for i := 32 ;i <= 126; i++{
// 	//1 	

// }


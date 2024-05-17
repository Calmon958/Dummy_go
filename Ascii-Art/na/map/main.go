package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter message to be printed")
		return
	}

	s, err := os.ReadFile("../../standard.txt")
	if err != nil {
		fmt.Println("text file not found")
		return
	}

	m := asciiArt(string(s))
	Print(strings.Join(os.Args[1:], " "), m)
}

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

// for loop - grouping each ascii art represented in the 8 lines
func PrintLine(text string, m map[rune][]string) {
	for b := 0; b < 8; b++ {
		for _, value := range text {
			fmt.Print(m[value][b])
		}
		fmt.Println()
	}
}

func Print(paragraph string, m map[rune][]string) {
	c := strings.Split(paragraph, "\\n")
	for _, line := range c {
		if line == "" {
			fmt.Println()
		} else {
			PrintLine(line, m)
		}
	}
}

// represents the posibility of handling extra lines found in the message to be printed

func Tab(s string) string {
	// word := strings.Split(string(s), "\n")
	j := len(s)
	result := ""
	for i := 0; i < j; i++ {
		// loops through each rune of a given message
		if s[i] == '\\' && s[i+1] == 't' {
			// checks if the character we are looking for is available
			result = result + "  "
			// replaces the \ with double space
			i++
			// moves to the next rune
		} else {
			result += string(s[i])
			// returns the rune if the character doesn't have the desired character
		}
	}
	return result
}

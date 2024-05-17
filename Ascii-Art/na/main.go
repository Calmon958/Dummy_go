package main

import (
	"os"
	"log"
	"strings"
)

func main() {
	asciiTxt, err := os.ReadFile("../standard.txt") // where the ascii text is stored
	if err != nil {
		log.Fatalf("File is empty. Please choose another file")
	}
	wordArr := strings.Split(string(asciiTxt), "\n")
	l := len(wordArr)
	charCounter := ' '

	m := make(map[rune][]string)
	var buffer []string

	for i := 0; i < l; i++ {
		line := wordArr[i]

		if i == 0 {
			buffer = make([]string, 0)
		} else if i % 9 == 0 {
			m[charCounter] = buffer
			charCounter++
			buffer = make([]string, 0)
		} else {
			buffer = append(buffer, line)

		}
	}

	
}
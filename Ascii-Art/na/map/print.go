package main

import (
	"fmt"
	"strings"
)

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

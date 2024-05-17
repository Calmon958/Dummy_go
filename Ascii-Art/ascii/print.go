package main

import (
	"fmt"
	"strings"
)

func Paragraph(sent string, m map[rune][]string) {
	if sent == "" {
		return
	}
	i := 0
	c := strings.Split(sent, "\\n")
	for _, line := range c {
		// fmt.Printf("Line: %s\n", line)
		if line == "" {
			i++
			if i < len(c) {
				fmt.Println()
			}
		} else {
			PrintLine(line, m)
		}
	}
}

// represents the posibility of handling extra lines found in the message to be printed

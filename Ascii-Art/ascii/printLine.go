package main

import (
	"fmt"
	"os"
)

func (text string, m map[rune][]string) {
	s := ""
	for b := 0; b < 8; b++ {
		for _, value := range text {
			mv, ok := m[value]
			if !ok {
				fmt.Println("Non-ASCII character")
				os.Exit(1)
			}
			// fmt.Print(mv[b])
			s += mv[b]
		}
		// fmt.Println()
		s += "\n"

	}

	fmt.Print(s)
}

package main

import "fmt"

func PrintLine(text string, m map[rune][]string) {
	for b := 0; b < 8; b++ {
		for _, value := range text {
			fmt.Print(m[value][b])
		}
		fmt.Println()
	}
}

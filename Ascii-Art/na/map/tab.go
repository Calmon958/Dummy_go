package main

import "regexp"


func Tab(s string) string {
	// 	// word := strings.Split(string(s), "\n")
	// 	j := len(s)
	
	//		for i := 0; i < j; i++ {
	//			// loops through each rune of a given message
	//			if s[i] == '\\' && s[i+1] == 't' {
	//				// checks if the character we are looking for is available
	//				result = result + "  "
	//				// replaces the \ with double space
	//				i++
	//				// moves to the next rune
	//			} else {
	//				result += string(s[i])
	//				// returns the rune if the character doesn't have the desired character
	//			}
	//		}
	//		return result
	//	}
	result := ""

	var feature string
	feature = `(/s*)(\\t+)(\s*)`
	result = regexp.MustCompile(feature).ReplaceAllString(s, "  ")
	return result
}

func Space(s string) string {
	result := ""
	var feature2 string
	var feature2 := `\s+(" ")\s+`
	result := 
}

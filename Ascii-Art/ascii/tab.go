package main

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


for _, value := range s {


package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: expected go run . <inputFile> <outputFile>")
		// reading arguments from the command line
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	// reading contents of the source file in bytes
	// input, err := os.ReadFile(inputFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	array := string(data)
	// converts the content from bytes to string.


	result := operation(array) + "\n"
	outputData := []byte(result)
	err = os.WriteFile(outputFile, outputData, 0644)
	if err != nil {
		panic(err)
	}
	// outFile := outputFile



}
func operation(array string) string {

	strArray := strings.Split(array, " ")
	for a, words := range strArray {
		if words == "(hex)" {
			// if there is a slice with (hex) it changes the word before it to int64
			hexDec, err := strconv.ParseInt(strArray[a-1], 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			strArray[a-1] = strconv.Itoa(int(hexDec))
			// changes the result from int64 to binary
		} else if words == "(bin)" {
			// changes the word before (bin) from binary to int64
			binDec, err := strconv.ParseInt(strArray[a-1], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			strArray[a-1] = strconv.Itoa(int(binDec))
			// converts the int64 to decimal
		} else if words == "(up)" {
			// takes the word before (up) and converts it to uppercase
			strArray[a-1] = strings.ToUpper(strArray[a-1])
		} else if words == "(cap)" {
			// takes the word before (cap)and coverts the first letter to caps
			strArray[a-1] = strings.Title(strArray[a-1])
		} else if words == "(low)" {
			// takes the word before (low) and converts it to lowercase
			strArray[a-1] = strings.ToLower(strArray[a-1])
		} else if words == "(cap," {
			series := strings.Trim(strArray[a+1], strArray[a+1][1:])
			// Takes the word that matches "(cap, and as the index position and moves to the next word after it taking
			// the first character alone.
			num, err := strconv.Atoi(series)
			if err != nil {
				log.Fatal(err)
			}
			// converts the string obtained into an int
			// strArray'[a-1] = strings.ToUpper(num)
			for b := 1; b <= num; b++ {
				strArray[a-b] = strings.Title(strArray[a-b])
			}
			// takes the previous words before the caps and capitalizes it

		} else if words == "(low," {
			steps := strings.Trim(strArray[a+1], strArray[a+1][1:])
			// Takes the word that matches "(cap, and as the index position and moves to the next word after it taking
			// the first character alone.
			num, err := strconv.Atoi(steps)
			if err != nil {
				log.Fatal(err)
			}
			// converts the string obtained into an int
			// strArray'[a-1] = strings.ToUpper(num)
			for b := 1; b <= num; b++ {
				strArray[a-b] = strings.ToLower(strArray[a-b])
			}
			// takes the previous words before the caps and capitalizes it

		} else if words == "(up," {
			steps := strings.Trim(strArray[a+1], strArray[a+1][1:])
			// Takes the word that matches "(cap, and as the index position and moves to the next word after it taking
			// the first character alone.
			num, err := strconv.Atoi(steps)
			if err != nil {
				log.Fatal(err)
			}
			// converts the string obtained into an int
			// strArray'[a-1] = strings.ToUpper(num)
			for b := 1; b <= num; b++ {
				strArray[a-b] = strings.ToUpper(strArray[a-b])
			}
			// takes the previous words before the caps and capitalizes it

		} else {

			// stored all the required letter in an array so that I can loop through them
			// punc := strings.Trim(strArray')

			letters := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
			for _, letter := range letters {
				if strArray[a] == "a" && string(strArray[a+1][0]) == letter {
					strArray[a] = "an"
				} else if strArray[a] == "A" && string(strArray[a+1][0]) == letter {
					strArray[a] = "An"
				}
			}
		}
	}

	// res, err := os.Create(outputFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	sentence := strings.Join(strArray, " ")

	removal := `\((hex|bin|cap|low|up)(,\s*\d+)?\)\s*`
	sentence = regexp.MustCompile(removal).ReplaceAllString(sentence, "")

	// regular expression for replacing cap, hex, bin, up, low words that were instructions from the output

	punctuations := `(.*?)(\s+)([.:;,?!])`
	sentence = regexp.MustCompile(punctuations).ReplaceAllString(sentence, "$1$3 ")

	// regular expression for removing spaces before punctuations

	punctuations_2 := `([,.:;?!])(\s+)([,.:;?!])`
	sentence = regexp.MustCompile(punctuations_2).ReplaceAllString(sentence, "$1$3")

	apostrophe := `(\s*)(')(\s*)`
	sentence = regexp.MustCompile(apostrophe).ReplaceAllString(sentence, "$2")

	// result, err := io.WriteString(res, sentence)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return sentence
}

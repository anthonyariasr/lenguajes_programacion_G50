package main

import (
	"strconv"
	"strings"
)

func CountWordsandLetters(text string, textDescrip string) string {
	lineJumps := strings.Count(text, "\n")
	spaces := strings.Count(text, " ")
	numChars := (len(text) - spaces - lineJumps)

	returnStr := "The " + textDescrip + " has " + strconv.Itoa(numChars) + " characters, " + strconv.Itoa(spaces+1) + " words, and " + strconv.Itoa(lineJumps+1) + " lines"
	return (returnStr)

}

func RotateSequence() {

}

func ShoesInventory() {

}

func main() {

	//Exercise 1

	text := `The late twentieth century has witnessed a scientific gold rush of astonishing proportions: the 
headlong and furious haste to commercialize genetic engineering. This enterprise has proceeded 
so rapidly-with so little outside commentary-that its dimensions and implications are hardly 
understood at all.`

	textDecription := "first paragraph of Michael Crichton's Jurassic Park"

	result := CountWordsandLetters(text, textDecription)

	print(result)

	//Exercise 3
	RotateSequence()

	//Exercise 4
	ShoesInventory()
}

package main

import (
	"fmt"
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

func remove(sliceArray []string, s int) []string {
	return append(sliceArray[:s], sliceArray[s+1:]...)
}

func RotateSequence(rotaQua int, direction bool, array *[]string) {

	//true == right
	//false == left

	if direction == true {
		for i := 0; i < rotaQua; i++ {
			*array = append([]string{(*array)[len(*array)-1]}, *array...)
			*array = remove(*array, len(*array)-1)
		}

	} else {
		for i := 0; i < rotaQua; i++ {
			*array = append(*array, (*array)[0])
			*array = remove(*array, 0)
		}
	}
}

func ShoesInventory() {

}

func main() {

	//Exercise 1
	print("\n\nExercise 1: \n")
	text := `The late twentieth century has witnessed a scientific gold rush of astonishing proportions: the 
headlong and furious haste to commercialize genetic engineering. This enterprise has proceeded 
so rapidly-with so little outside commentary-that its dimensions and implications are hardly 
understood at all.`

	textDecription := "first paragraph of Michael Crichton's Jurassic Park"

	result := CountWordsandLetters(text, textDecription)

	print(result)

	//Exercise 3
	print("\n\nExercise 3: \n")

	//Amount of times the elements will rotate
	numPos := 3

	//true == right
	//false == left
	direct := true

	originalArray := &[]string{"a", "b", "c", "d", "e"}

	print("The original Array is: ")
	fmt.Println(*originalArray)

	RotateSequence(numPos, direct, originalArray)

	fmt.Print("The array rotated " + strconv.Itoa(numPos) + " times is: ")
	fmt.Println(*originalArray)

	//Exercise 4
	ShoesInventory()
}

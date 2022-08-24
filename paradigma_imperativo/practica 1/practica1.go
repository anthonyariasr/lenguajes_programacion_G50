package main

import (
	"fmt"
	"strconv"
	"strings"
)

//______________________________________________________________________________________________________________________________________________________________________

func CountWordsandLetters(text string, textDescrip string) string {
	lineJumps := strings.Count(text, "\n")
	spaces := strings.Count(text, " ")
	numChars := (len(text) - spaces - lineJumps)

	returnStr := "The " + textDescrip + " has " + strconv.Itoa(numChars) + " characters, " + strconv.Itoa(spaces+1) + " words, and " + strconv.Itoa(lineJumps+1) + " lines"
	return (returnStr)

}

//______________________________________________________________________________________________________________________________________________________________________

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

//______________________________________________________________________________________________________________________________________________________________________

type calzado struct {
	marca     string
	talla     int
	precio    int
	cantStock int
}

var arrayCalzado []calzado

// validates if the shoe has the right size and then adds it into the array
func addShoe(shoe calzado) {
	if shoe.talla >= 34 && shoe.talla <= 44 {
		arrayCalzado = append(arrayCalzado, shoe)
	} else {
		print("The shoe size must be between 34 and 44")
	}
}

// Search for the shoe and then check if there's any in stock
func sellShoe(Pmarca string) {
	var index = -1
	for i, shoe := range arrayCalzado {
		if shoe.marca == Pmarca {
			index = i
		}
	}

	if index != -1 && arrayCalzado[index].cantStock > 0 {
		arrayCalzado[index].cantStock--
		fmt.Println("Shoe has been sold")
	} else {
		fmt.Println("This shoe can't be find in stock")
	}
}

func ShoesInventory() {

	addShoe(calzado{marca: "Adidas", talla: 42, precio: 50000, cantStock: 5})
	addShoe(calzado{marca: "Nike", talla: 44, precio: 250000, cantStock: 1})
	addShoe(calzado{marca: "Rebook", talla: 36, precio: 125000, cantStock: 2})
	fmt.Println(arrayCalzado)
	sellShoe("Adidas")
	fmt.Println(arrayCalzado)
}

//______________________________________________________________________________________________________________________________________________________________________

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
	fmt.Print("\n\nExercise 4: \n")
	ShoesInventory()
}

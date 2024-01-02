package exercise_1

import (
	"fmt"
)

func FirstExercise(word string) {
	var numberOfLetters uint8
	// Ejercicio #1: letras de una palabra
	for _, letter := range word {
		numberOfLetters++
		// Forma 1:
		fmt.Printf("Letra: %c \n", letter)
		// Forma 2:
		//fmt.Printf("Letra: %s \n", string(letter))
	}

	fmt.Printf("La palabra %s tiene %v letras \n", word, numberOfLetters)

}

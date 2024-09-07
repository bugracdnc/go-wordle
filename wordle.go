package main

import (
	"fmt"
	"math/rand/v2"
)

var wordleDict = []string{"BUGRA", "NESLI"}

func main() {
	randSelected := wordleDict[rand.IntN(len(wordleDict))]

	var userInput string
	errString := ""
	var tries = 6

	fmt.Println("_ _ _ _ _\n")

	isGuessed := false
	for i := 0; i < tries; i++ {
		for len(userInput) != 5 {
			fmt.Printf("%d tries left\n", (tries - i))
			fmt.Printf("Enter a %sguess: ", errString)
			fmt.Scanln(&userInput)
			errString = "five length "
		}

		fmt.Println()
		isCorrect := true

		for index, _ := range randSelected {
			if randSelected[index] == userInput[index] {
				fmt.Print(string(randSelected[index]))
			} else {
				isCorrect = false
				fmt.Print("_")
			}
			fmt.Print(" ")
		}
		fmt.Println("\n")
		if isCorrect {
			isGuessed = true
			break
		} else {
			userInput = ""
			errString = ""
		}
	}
	if isGuessed {
		fmt.Println("Congratulations!\nYou win!\n\n")
	} else {
		fmt.Println("Sorry, you run out of tries.\n\n")
	}
}

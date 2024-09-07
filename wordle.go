package main

import (
	"fmt"
	"math/rand/v2"
)

type Color struct {
	Reset   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	Gray    string
	White   string
}

func initColor() *Color {
	var color Color
	color.Reset = "\033[0m"
	color.Red = "\033[31m"
	color.Green = "\033[32m"
	color.Yellow = "\033[33m"
	color.Blue = "\033[34m"
	color.Magenta = "\033[35m"
	color.Cyan = "\033[36m"
	color.Gray = "\033[37m"
	color.White = "\033[97m"

	return &color
}

var wordleDict = []string{"BUGRA", "NESLI"}

func main() {
	color := initColor()
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
				fmt.Print(color.Green + string(randSelected[index]) + color.Reset)
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

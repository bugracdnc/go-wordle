package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const totalTries int = 6
const wordLen int = 5

type Wordle struct {
	index    int
	dict     WordleDict
	word     string
	tries    int
	solved   [wordLen]byte
	attempts [totalTries]string
}

func initWordle() *Wordle {
	var w Wordle
	w.dict = *initWordleDict()
	w.index = rand.IntN(len(w.dict.elems))
	w.word = w.dict.elems[w.index]
	w.tries = totalTries
	w.solved = [wordLen]byte{0}
	return &w
}

func parseUserInput(wordle *Wordle, userInput string) bool {
	var builder strings.Builder
	color := initColor()
	toFind := make(map[byte]int)
	for _, c := range []byte(wordle.word) {
		toFind[c]++
	}
	found := make(map[byte]int)
	foundc := 0

	//scan through the source string for input
	for i := 0; i < wordLen; i++ {
		inputChar := userInput[i]
		srcChar := wordle.word[i]
		if inputChar == srcChar {
			foundc++
			wordle.solved[i] = srcChar
			toFind[srcChar]--
		}
	}

	//scan again for correct values in wrong places
	if foundc < wordLen {
		for i := 0; i < wordLen; i++ {
			if wordle.solved[i] == 0 {
				inputChar := userInput[i]
				if count(wordle.word, inputChar) > 0 && toFind[inputChar] > 0 {
					found[inputChar]++
					toFind[inputChar]--
				}
			}
		}
	}

	//build result for user
	for i := 0; i < wordLen; i++ {
		if userInput[i] == wordle.solved[i] {
			builder.WriteString(color.Green + string(wordle.solved[i]) + color.Reset)
		} else if found[userInput[i]] > 0 {
			builder.WriteString(color.Yellow + string(userInput[i]) + color.Reset)
			found[userInput[i]]--
		} else {
			builder.WriteByte(userInput[i])
		}
		builder.WriteString(" ")
	}

	wordle.attempts[totalTries-wordle.tries] = builder.String()

	for i := 0; i < totalTries-wordle.tries+1; i++ {
		fmt.Println(wordle.attempts[i])
	}
	fmt.Println()
	return foundc == wordLen
}

func getUserInput(tries int) string {
	userInput := ""
	errString := ""
	for len(userInput) != wordLen {
		fmt.Printf("%d tries left\n", tries)
		fmt.Printf("Enter a %sguess: ", errString)
		fmt.Scanln(&userInput)
		errString = fmt.Sprintf("%d length ", wordLen)
	}
	fmt.Println()
	return userInput
}

func game(wordle *Wordle) bool {
	for ; wordle.tries > 0; wordle.tries-- {
		if parseUserInput(wordle, getUserInput(wordle.tries)) {
			return true
		}
	}
	return false
}

func main() {
	wordle := initWordle()
	fmt.Println()
	fmt.Println("### WORDLE ###")
	fmt.Println()
	fmt.Println("_ _ _ _ _")
	fmt.Println()

	if game(wordle) {
		fmt.Printf("Wordle #%d %d/%d\n\n", wordle.index+1, wordle.tries, totalTries)
		fmt.Println("Congratulations! You win!")
	} else {
		fmt.Println("Sorry, you run out of tries.")
	}
}

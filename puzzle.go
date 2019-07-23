package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	var str = "))(()()"
	numOfBracket := bracketCounter(str)
	fmt.Println("Number of valid bracket:", numOfBracket)

	permutate([]rune("abc"),0)
}

func bracketCounter(strWithBracket string) (bracketCounter int) {
	stk := stack.New();

	for _, elem := range strWithBracket {

		if elem == '(' {
			stk.Push(elem)

		} else if elem == ')' {
			if stk.Pop() != nil {
				bracketCounter++
			}
		}
	}
	return
}

var looping_counter = 0

func printRune(char_arr []rune) {
	fmt.Println(looping_counter, string(char_arr))
}

func permutate(word []rune, i int) {

	if i > len(word) {
		printRune(word)
		return
	}

	pos := i + 1;

	permutate(word, pos)

	if (pos > len(word)) {
		return
	}

	for j := i + 1; j < len(word); j++ {
		looping_counter++
		word[i], word[j] = word[j], word[i]
		permutate(word, i + 1)
		word[i], word[j] = word[j], word[i]
	}
}


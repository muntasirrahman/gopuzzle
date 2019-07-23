## Valid Bracket Counter

### How it work
Put every opening bracket into a stack, and pop it for every closing bracket
```go
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
```

## Permutation Generator

### How it works
pick one characther, permutate the remaining char array

### Output
- abc
- acb
- bac
- bca
- cba
- cab

```go
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

```
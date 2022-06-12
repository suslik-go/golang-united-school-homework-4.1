package reverse_string

func ReverseString(input string) (output string) {

	runesList := []rune(input)

	for i := len(runesList) - 1; i >= 0; i-- {
		output += string(runesList[i])
	}

	return output
}

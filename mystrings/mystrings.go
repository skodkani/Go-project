package mystrings

// Reverse the string from left to right.
// The first character of the function should be in Capital, So that
// the func is accessible when the mystring package is imported.
func Reverse(s string) string {
	// Storing the reversed string
	result := ""

	for _, v := range s {
		// Need to convert the char to string, since its rune not a character.
		// Concatenates the Char to the result.
		result = string(v) + result
	}

	// Return the Reversed String
	return result
}

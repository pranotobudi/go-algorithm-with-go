package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {

	chars := strings.Split(word, "")
	result := []string{}
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}

	return strings.Join(result, "")
}

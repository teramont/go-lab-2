package lab2

import (
	"fmt"
	"regexp"
	"strings"
)

var r = regexp.MustCompile("\\+|\\*|\\/|\\-|\\^|[0-9]*")

func splitIntoTokens(input string) []string {
	r := strings.NewReplacer(
		"+", " + ",
		"-", " - ",
		"/", " / ",
		"*", " * ",
		"^", " ^ ",
	)

	// Adding white characters to simplify token splitting
	withWhitespaces := r.Replace(input)
	tokens := strings.Fields(withWhitespaces)
	return tokens
}

// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}

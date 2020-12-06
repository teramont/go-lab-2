package lab2

import (
	"fmt"
	"strings"
)

type stack struct {
	stack []string
}

func (s *stack) push(str string) {
	s.stack = append(s.stack, str)
}

func (s *stack) len() int {
	return len(s.stack)
}

func (s *stack) pop() *string {
	if len(s.stack) == 0 {
		return nil
	}
	lastItem := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return &lastItem
}

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

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}

func grammarError(s string) error {
	return fmt.Errorf("Grammar error after '%s'", s)
}

// PostfixToPrefix converts postfix notation into the prefix one
func PostfixToPrefix(input string) (string, error) {
	tokens := splitIntoTokens(input)
	stack := stack{}
	for _, s := range tokens {
		if isOperator(s) {
			a := stack.pop()
			b := stack.pop()

			if a == nil {
				return "", grammarError(s)
			} else if b == nil {
				return "", grammarError(*a)
			}
			stack.push(fmt.Sprintf("%s %s %s", s, *b, *a))
		} else {
			stack.push(s)
		}
	}
	if stack.len() > 1 {
		return "", fmt.Errorf("Multiple expressions")
	}
	res := stack.pop()
	return *res, nil
}

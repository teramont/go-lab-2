package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"18", "18"},
		{"2 2 +", "+ 2 2"},
		{"4 2 - 3 * var +", "+ * - 4 2 3 var"},
		{"4 2-3*юнікод+", "+ * - 4 2 3 юнікод"},
	}

	for _, c := range cases {
		res, err := PostfixToPrefix(c.in)
		if assert.Nil(t, err) {
			assert.Equal(t, c.want, res)
		}
	}

	casesWrong := []string{
		"+ 1",
		"1 +",
		"2 2 + 2 2 +",
	}

	for _, c := range casesWrong {
		_, err := PostfixToPrefix(c)
		if !assert.NotNil(t, err) {
			t.Fatal("Not Nil")
		}
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PostfixToPrefix("2 2 +")
	fmt.Println(res)

	// Output:
	// + 2 2
}

func TestTokenizer(t *testing.T) {
	res := splitIntoTokens("  123+456 /   978 ^variable-юнікод*  ")
	assert.Equal(t, res, []string{"123", "+", "456", "/", "978", "^", "variable", "-", "юнікод", "*"})
}

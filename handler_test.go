package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	r := strings.NewReader("2 2 +")
	var w bytes.Buffer

	handler := ComputeHandler{Input: r, Output: &w}
	err := handler.Compute()
	assert.Nil(t, err)
	got := w.String()
	assert.Equal(t, got, "+ 2 2\n")
}

func TestErrorCompute(t *testing.T) {
	r := strings.NewReader("2 +")
	var w bytes.Buffer

	handler := ComputeHandler{Input: r, Output: &w}
	err := handler.Compute()
	assert.NotNil(t, err)
	got := w.String()
	assert.Equal(t, got, "")
}

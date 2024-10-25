package stu

import (
	"reflect"
	"testing"
)

// Test for the `tokenize` function.
func TestTokenize(t *testing.T) {
	template := `Hello, {{ name }}!`

	expectedTokens := []Token{
		{Type: TextToken, Literal: "Hello, "},
		{Type: VarToken, Literal: "name"},
		{Type: TextToken, Literal: "!"},
	}

	tokens := tokenize(template)

	if !reflect.DeepEqual(tokens, expectedTokens) {
		t.Errorf("Tokenize failed. Expected %v, got %v", expectedTokens, tokens)
	}
}

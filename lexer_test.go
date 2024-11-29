package stu

import "testing"

func TestTokenize(t *testing.T) {
	tests := []struct {
		template string
		expected []Token
	}{
		{
			"Hello, {{ name }}!",
			[]Token{
				{Type: TOKEN_TEXT, Value: "Hello, ", Line: 1, Col: 1},
				{Type: TOKEN_VARIABLE, Value: "name", Line: 1, Col: 8},
				{Type: TOKEN_TEXT, Value: "!", Line: 1, Col: 13},
				{Type: TOKEN_EOF, Value: "", Line: 1, Col: 14},
			},
		},
		{
			"{% for item in items %}{{ item }}{% endfor %}",
			[]Token{
				{Type: TOKEN_FOR, Value: "for item in items", Line: 1, Col: 1},
				{Type: TOKEN_VARIABLE, Value: "item", Line: 1, Col: 19},
				{Type: TOKEN_END_FOR, Value: "endfor", Line: 1, Col: 26},
				{Type: TOKEN_EOF, Value: "", Line: 1, Col: 33},
			},
		},
		{
			"{# This is a comment #}",
			[]Token{
				{Type: TOKEN_COMMENT, Value: "This is a comment", Line: 1, Col: 1},
				{Type: TOKEN_EOF, Value: "", Line: 1, Col: 22},
			},
		},
		{
			`<h1>Welcome, {{ name }}!</h1>

<p>This is the first text element.</p>

<ul>
  {% for item in items %}
    <li>{{ item }}</li>
  {% endfor %}
</ul>

<p>This is the second text element.</p>`,
			[]Token{
				{Type: TOKEN_TEXT, Value: "<h1>Welcome, ", Line: 1, Col: 1},
				{Type: TOKEN_VARIABLE, Value: "name", Line: 1, Col: 14},
				{Type: TOKEN_TEXT, Value: "!</h1>\n\n<p>This is the first text element.</p>\n\n<ul>\n  ", Line: 1, Col: 19},
				{Type: TOKEN_FOR, Value: "for item in items", Line: 5, Col: 3},
				{Type: TOKEN_TEXT, Value: "\n    <li>", Line: 6, Col: 5},
				{Type: TOKEN_VARIABLE, Value: "item", Line: 6, Col: 5},
				{Type: TOKEN_TEXT, Value: "</li>\n  ", Line: 7, Col: 6},
				{Type: TOKEN_END_FOR, Value: "endfor", Line: 8, Col: 5},
				{Type: TOKEN_TEXT, Value: "\n</ul>\n\n<p>This is the second text element.</p>", Line: 9, Col: 1},
				{Type: TOKEN_EOF, Value: "", Line: 10, Col: 56},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.template, func(t *testing.T) {
			tokens := tokenize(test.template)
			for i, token := range tokens {
				if token.Type != test.expected[i].Type || token.Value != test.expected[i].Value {
					t.Errorf("expected %v, got %v", test.expected[i], token)
				}
			}
		})
	}
}

package stu

import "strings"

type TokenType int

const (
	TextToken  TokenType = iota
	VarToken             // {{ }}
	BlockToken           // {% %}
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}

func tokenize(template string) []Token {
	var tokens []Token
	var buffer string
	var insideVar, insideBlock bool

	for _, char := range template {
		buffer += string(char)

		// Handle start of variable `{{ ... }}`
		if strings.HasSuffix(buffer, "{{") && !insideVar && !insideBlock {
			if len(buffer) > 2 {
				tokens = append(tokens, Token{Type: TextToken, Literal: buffer[:len(buffer)-2]})
			}
			insideVar = true
			buffer = "{{"
		} else if strings.HasSuffix(buffer, "}}") && insideVar {
			tokens = append(tokens, Token{Type: VarToken, Literal: strings.TrimSpace(buffer[2 : len(buffer)-2])})
			insideVar = false
			buffer = ""

			// Handle start of block `{% ... %}`
		} else if strings.HasSuffix(buffer, "{%") && !insideVar && !insideBlock {
			if len(buffer) > 2 {
				tokens = append(tokens, Token{Type: TextToken, Literal: buffer[:len(buffer)-2]})
			}
			insideBlock = true
			buffer = "{%"
		} else if strings.HasSuffix(buffer, "%}") && insideBlock {
			tokens = append(tokens, Token{Type: BlockToken, Literal: strings.TrimSpace(buffer[2 : len(buffer)-2])})
			insideBlock = false
			buffer = ""
		}
	}

	// Append any remaining text as a TextToken
	if buffer != "" {
		tokens = append(tokens, Token{Type: TextToken, Literal: buffer})
	}

	return tokens
}

// Parse function - takes tokens, returns nodes (AST)
func parseTemplate(tokens []Token) []Node {
	var nodes []Node

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type {
		case TextToken:
			nodes = append(nodes, &TextNode{Value: token.Literal}) // Append a text node
		case VarToken:
			nodes = append(nodes, &VarNode{Name: token.Literal}) // Append a variable node
		case BlockToken:
			// Handling block tokens like {% if ... %}, {% for ... %} could be more complex,
			// we'll skip detailed processing here for simplicity, but you'd need to handle
			// full conditional and loop logic here.
		case EOF:
			// End of token stream
			break
		}
	}

	return nodes
}

package stu

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

	insideVar, insideBlock := false, false

	for _, char := range template {
		buffer += string(char)

		if buffer == "{{" {
			tokens = append(tokens, Token{Type: TextToken, Literal: buffer[:len(buffer)-2]})
			insideVar = true
			buffer = "{{"
		} else if buffer == "}}" && insideVar {
			tokens = append(tokens, Token{Type: VarToken, Literal: buffer})
			insideVar = false
			buffer = ""
		} else if buffer == "{%" {
			tokens = append(tokens, Token{Type: TextToken, Literal: buffer[:len(buffer)-2]})
			insideBlock = true
			buffer = "{%"
		} else if buffer == "%}" && insideBlock {
			tokens = append(tokens, Token{Type: BlockToken, Literal: buffer})
			insideBlock = false
			buffer = ""
		}
	}

	if buffer != "" {
		tokens = append(tokens, Token{Type: TextToken, Literal: buffer})
	}

	return tokens
}

// Parse function - takes tokens, returns nodes (AST)
func parseTemplate(tokens []Token) []Node {
	var nodes []Node

	for i := 0; i < len(tokens); i++ {
		//token := tokens[i]
		//switch token.Type {
		//case TextToken:
		//	nodes = append(nodes, &TextNode{Text: token.Literal}) // Append a text node
		//case VarToken:
		//	// Variable tokens like {{ name }} - we need to trim the curly braces and spaces
		//	varName := strings.TrimSpace(token.Literal[2 : len(token.Literal)-2]) // Remove `{{` and `}}`
		//	nodes = append(nodes, &VarNode{VarName: varName})                     // Append a variable node
		//case BlockToken:
		//	// Handling block tokens like {% if ... %}, {% for ... %} could be more complex,
		//	// we'll skip detailed processing here for simplicity, but you'd need to handle
		//	// full conditional and loop logic here.
		//case EOFToken:
		//	// End of token stream
		//	break
		//}
	}

	return nodes
}

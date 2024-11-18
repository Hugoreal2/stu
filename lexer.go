package stu

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	TOKEN_TEXT      TokenType = iota // 0: Plain text
	TOKEN_VARIABLE                   // 1: {{ variable }}
	TOKEN_BLOCK                      // 2: {% block %}
	TOKEN_END_BLOCK                  // 3: {% endblock %}
	TOKEN_FOR                        // 4: {% for %}
	TOKEN_END_FOR                    // 5: {% endfor %}
	TOKEN_IF                         // 6: {% if %}
	TOKEN_END_IF                     // 7: {% endif %}
	TOKEN_ELSE                       // 8: {% else %}
	TOKEN_COMMENT                    // 9: {# comment #}
	TOKEN_EOF                        // 10: End of file
)

type Token struct {
	Type  TokenType // The type of token (e.g., TEXT, VARIABLE)
	Value string    // The actual content of the token
	Line  int       // Line number for error reporting
	Col   int       // Column number for error reporting
}

func tokenize(template string) []Token {
	var tokens []Token
	length := len(template)
	line, col := 1, 1

	for i := 0; i < length; {
		// Check for variable: {{ variable }}
		if strings.HasPrefix(template[i:], "{{") {
			end := strings.Index(template[i:], "}}")
			if end == -1 {
				panic("Unclosed variable tag")
			}
			value := strings.TrimSpace(template[i+2 : i+end])
			tokens = append(tokens, Token{Type: TOKEN_VARIABLE, Value: value, Line: line, Col: col})
			i += end + 2
			col += end + 2
		} else if strings.HasPrefix(template[i:], "{%") {
			// Check for blocks like {% block %}, {% for %}, {% if %}
			end := strings.Index(template[i:], "%}")
			if end == -1 {
				panic("Unclosed block tag")
			}

			tagContent := strings.TrimSpace(template[i+2 : i+end])
			switch {
			case strings.HasPrefix(tagContent, "for"):
				tokens = append(tokens, Token{Type: TOKEN_FOR, Value: tagContent, Line: line, Col: col})
			case strings.HasPrefix(tagContent, "endfor"):
				tokens = append(tokens, Token{Type: TOKEN_END_FOR, Value: tagContent, Line: line, Col: col})
			case strings.HasPrefix(tagContent, "if"):
				tokens = append(tokens, Token{Type: TOKEN_IF, Value: tagContent, Line: line, Col: col})
			case strings.HasPrefix(tagContent, "endif"):
				tokens = append(tokens, Token{Type: TOKEN_END_IF, Value: tagContent, Line: line, Col: col})
			case strings.HasPrefix(tagContent, "block"):
				tokens = append(tokens, Token{Type: TOKEN_BLOCK, Value: tagContent, Line: line, Col: col})
			case strings.HasPrefix(tagContent, "endblock"):
				tokens = append(tokens, Token{Type: TOKEN_END_BLOCK, Value: tagContent, Line: line, Col: col})
			case strings.HasPrefix(tagContent, "else"):
				tokens = append(tokens, Token{Type: TOKEN_ELSE, Value: tagContent, Line: line, Col: col})
			default:
				// Handle other block types or invalid block tags
				tokens = append(tokens, Token{Type: TOKEN_BLOCK, Value: tagContent, Line: line, Col: col})
			}

			// Move the index forward
			i += end + 2
			col += end + 2
		} else if strings.HasPrefix(template[i:], "{#") {
			// Check for comment: {# comment #}
			end := strings.Index(template[i:], "#}")
			if end == -1 {
				panic("Unclosed comment tag")
			}
			value := strings.TrimSpace(template[i+2 : i+end])
			tokens = append(tokens, Token{Type: TOKEN_COMMENT, Value: value, Line: line, Col: col})
			i += end + 2
			col += end + 2
		} else {
			// Handle plain text
			end := i
			for end < length && !strings.HasPrefix(template[end:], "{{") && !strings.HasPrefix(template[end:], "{%") && !strings.HasPrefix(template[end:], "{#") {
				if template[end] == '\n' {
					line++
					col = 0
				}
				end++
				col++
			}
			if end > i {
				value := template[i:end]
				tokens = append(tokens, Token{Type: TOKEN_TEXT, Value: value, Line: line, Col: col})
				i = end
			}
		}
	}

	// Append EOF token
	tokens = append(tokens, Token{Type: TOKEN_EOF, Value: "", Line: line, Col: col})
	return tokens
}

// Parse function - takes tokens, returns nodes (AST)
func parseTemplate(tokens []Token) []Node {
	var nodes []Node
	length := len(tokens)
	i := 0

	for i < length {
		token := tokens[i]

		switch token.Type {
		case TOKEN_TEXT:
			// Create a TextNode
			nodes = append(nodes, &TextNode{Value: token.Value})

		case TOKEN_VARIABLE:
			// Create a VarNode
			nodes = append(nodes, &VarNode{Name: token.Value})

		case TOKEN_BLOCK:
			// Handle "for" and "endfor" blocks
			if token.Value == "for" {
				// Expect a loop definition (e.g., "item in items")
				i++
				if i >= length {
					panic("Unexpected end of tokens inside 'for' block")
				}
				loopDef := tokens[i].Value
				parts := splitLoopDefinition(loopDef)
				if len(parts) != 2 {
					panic("Invalid 'for' loop definition")
				}
				loopVarName := parts[0]
				iterList := parts[1]

				// Parse the body of the "for" loop
				var body []Node
				i++
				for i < length && tokens[i].Type != TOKEN_END_FOR {
					body = append(body, parseNode(tokens, &i))
				}

				// Create a ForNode
				nodes = append(nodes, &ForNode{
					LoopVarName: loopVarName,
					IterList:    iterList,
					Body:        body,
				})
			}
		case TOKEN_COMMENT:
			// Ignore comments
			i++
			continue

		case TOKEN_EOF:
			return nodes

		default:
			panic(fmt.Sprintf("Unexpected token type: %v", token.Type))
		}

		i++
	}

	return nodes
}

func parseNode(tokens []Token, i *int) Node {
	token := tokens[*i]
	switch token.Type {
	case TOKEN_TEXT:
		*i++
		return &TextNode{Value: token.Value}
	case TOKEN_VARIABLE:
		*i++
		return &VarNode{Name: token.Value}
	default:
		panic(fmt.Sprintf("Unexpected token type in parseNode: %v", token.Type))
	}
}

func splitLoopDefinition(loopDef string) []string {
	parts := strings.Fields(loopDef)
	if len(parts) != 3 || parts[1] != "in" {
		panic("Invalid loop definition. Expected format: 'item in items'")
	}
	return []string{parts[0], parts[2]}
}

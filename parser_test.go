package stu

import (
	"reflect"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	tokens := []Token{
		{Type: TextToken, Literal: "Hello, "},
		{Type: VarToken, Literal: "name"},
		{Type: TextToken, Literal: "!"},
	}

	expectedNodes := []Node{
		&TextNode{Value: "Hello, "},
		&VarNode{Name: "name"},
		&TextNode{Value: "!"},
	}

	nodes := parseTemplate(tokens)

	if !reflect.DeepEqual(nodes, expectedNodes) {
		t.Errorf("ParseTemplate failed. Expected %v, got %v", expectedNodes, nodes)
	}
}

func TestRenderTemplate(t *testing.T) {
	ast := []Node{
		&TextNode{Value: "Hello, "},
		&VarNode{Name: "name"},
		&TextNode{Value: "!"},
	}

	context := map[string]interface{}{
		"name": "Alice",
	}

	expectedOutput := "Hello, Alice!"

	output := renderTemplate(ast, context)

	if output != expectedOutput {
		t.Errorf("RenderTemplate failed. Expected %v, got %v", expectedOutput, output)
	}
}

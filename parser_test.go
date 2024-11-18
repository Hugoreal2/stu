package stu

import (
	"testing"
)

func TestParseTemplate(t *testing.T) {
	tests := []struct {
		template string
		expected []Node
	}{
		{
			"Hello, {{ name }}!",
			[]Node{
				&TextNode{Value: "Hello, "},
				&VarNode{Name: "name"},
				&TextNode{Value: "!"},
			},
		},
		{
			"{% for item in items %}{{ item }}{% endfor %}",
			[]Node{
				&ForNode{
					LoopVarName: "item",
					IterList:    "items",
					Body: []Node{
						&VarNode{Name: "item"},
					},
				},
			},
		},
		{
			"{# This is a comment #}",
			[]Node{}, // No nodes for comments, they're ignored in rendering
		},
	}

	for _, test := range tests {
		t.Run(test.template, func(t *testing.T) {
			tokens := tokenize(test.template)
			ast := parseTemplate(tokens)
			if len(ast) != len(test.expected) {
				t.Fatalf("expected %d nodes, got %d", len(test.expected), len(ast))
			}
			for i, node := range ast {
				switch n := node.(type) {
				case *TextNode:
					if tn, ok := test.expected[i].(*TextNode); ok && tn.Value != n.Value {
						t.Errorf("expected TextNode '%s', got '%s'", tn.Value, n.Value)
					}
				case *VarNode:
					if vn, ok := test.expected[i].(*VarNode); ok && vn.Name != n.Name {
						t.Errorf("expected VarNode '%s', got '%s'", vn.Name, n.Name)
					}
				case *ForNode:
					if fn, ok := test.expected[i].(*ForNode); ok {
						if fn.LoopVarName != n.LoopVarName || fn.IterList != n.IterList {
							t.Errorf("expected ForNode with loopVarName '%s' and iterList '%s', got '%s' and '%s'", fn.LoopVarName, fn.IterList, n.LoopVarName, n.IterList)
						}
					}
				}
			}
		})
	}
}

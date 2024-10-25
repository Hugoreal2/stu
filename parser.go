package stu

import (
	"fmt"
)

type Node interface {
	Render(context map[string]interface{}) string
}

type TextNode struct {
	Value string
}

func (t *TextNode) Render(ctx map[string]interface{}) string {
	return t.Value
}

type VarNode struct {
	Name string
}

func (v *VarNode) Render(ctx map[string]interface{}) string {
	value := ctx[v.Name]
	if value != nil {
		return fmt.Sprintf("%v", value)
	}
	return "Value Not Defined"
}

type ForNode struct {
	LoopVarName string
	IterList    string
	Body        []Node
}

func (f *ForNode) Render(ctx map[string]interface{}) string {
	// Assume ctx[f.IterList] is a list/array
	iterList, ok := ctx[f.IterList].([]interface{})
	if !ok {
		return ""
	}

	var result string
	for _, item := range iterList {
		innerCtx := map[string]interface{}{f.LoopVarName: item}
		for _, bodyNode := range f.Body {
			result += bodyNode.Render(innerCtx)
		}
	}

	return result
}

func renderTemplate(ast []Node, context map[string]interface{}) string {
	var result string
	for _, node := range ast {
		result += node.Render(context)
	}
	return result
}

package main

import (
	"fmt"

	"github.com/Hugoreal2/stu"
)

func main() {
	template := `Hello, {{ name }}!`
	context := map[string]interface{}{
		"name": "Alice",
	}
	ast := stu.ParseTemplate(stu.Tokenize(template))

	output := stu.RenderTemplate(ast, context)
	fmt.Println(output) // Hello, Alice!
}

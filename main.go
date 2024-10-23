package stu

import "fmt"

func main() {
	template := `Hello, {{ name }}!`
	context := map[string]interface{}{
		"name": "Alice",
	}
	ast := parseTemplate(tokenize(template))

	output := renderTemplate(ast, context)
	fmt.Println(output) // Hello, Alice!
}

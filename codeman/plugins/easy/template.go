package easy

import (
	"text/template"
)

func NewTextTmpl(name string, text string) *template.Template {
	t, err := template.New(name).Funcs(map[string]any{
		"a_a":   snakeCase,
		"A_A":   SnakeCase,
		"aA":    camelCase,
		"Aa":    CamelCase,
		"Ident": Ident,
	}).Parse(text)
	if err != nil {
		panic(err)
	}

	return t
}

package prototemplate

import (
	_ "embed"
	"fmt"
	"text/template"
)

//go:embed proto.tmpl
var protoTemplate string

func New(name string) *template.Template {
	tmpl := template.New(name)
	tmpl, err := template.New(name).Parse(protoTemplate)
	if err != nil {
		panic(fmt.Sprintf("failed to parse renderField template: %v", err))
	}
	return tmpl
}

package plugins

import (
	_ "embed"
	"text/template"
)

//go:embed templates/proto.tmpl
var protoTmpl string

func NewTemplate(name string) *template.Template {
	return template.New(name)
}

func ParseProtoTemplate(t *template.Template, content string) *template.Template {
	template.Must(t.Parse(protoTmpl))
	template.Must(t.Parse(content))
	return t
}

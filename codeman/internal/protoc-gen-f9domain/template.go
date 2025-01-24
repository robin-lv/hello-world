package main

import (
	_ "embed"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
	"tools/protobuf-plugins/easy"
	"tools/protobuf-plugins/f9options"
)

//go:embed server.tmpl
var serverTmpl string
var serverTemplate = easy.NewTextTmpl("server.tmpl", serverTmpl)

type (
	cmdDesc struct {
		Services []*serviceDesc
	}
	serviceDesc struct {
		Name     string
		BaseName string
		Methods  []*methodDesc
	}

	methodDesc struct {
		Name string
		CMD  string
	}
)

func buildServiceDesc(file *protogen.File) *cmdDesc {
	desc := &cmdDesc{}
	for _, svc := range file.Services {
		mod := &serviceDesc{}
		mod.Name = svc.GoName
		mod.BaseName = strings.TrimSuffix(mod.Name, "API")
		if mod.BaseName == "Base" {
			mod.BaseName = ""
		}
		for _, method := range svc.Methods {
			cmd := &methodDesc{Name: method.GoName, CMD: f9options.GetCMDOption(method)}
			mod.Methods = append(mod.Methods, cmd)
		}
		desc.Services = append(desc.Services, mod)
	}
	return desc
}

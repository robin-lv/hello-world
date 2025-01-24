package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"protoc-plugins/plugins/easy"
	"strings"
)

// version is the current protoc-gen-f9api version.
const version = "v1.0.0"
const pluginName = "protoc-gen-" + plugin
const plugin = "f9domain"

var (
	showVersion = flag.Bool("version", false, "print the version and exit")
	//omitempty       = flag.Bool("omitempty", true, "omit if google.api is empty")
	//omitemptyPrefix = flag.String("omitempty_prefix", "", "omit if google.api is empty")
	relativePath = flag.String("o", "", "out_dir : relative path")
	recovery     = flag.String("y", "exist", "recovery")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("%s %v\n", pluginName, version)
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {

		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}
			generateFile(gen, file)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	if len(file.Services) == 0 {
		return
	}
	d := buildServiceDesc(file)
	name := strings.TrimSuffix(d.Services[0].Name, "API")
	g := easy.NewGenFile2(gen, file, easy.Ident(name, "logic")+".go", pluginName, version)
	g.WriteGoFileHead(false)
	g.P()

	g.WriteTextTemplate(serverTemplate, d)
}

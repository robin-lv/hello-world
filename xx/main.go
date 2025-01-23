package main

import (
	_ "embed"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"protoc-plugins/plugins/protomodels"

	"protoc-plugins/plugins"
	"strings"
)

//go:embed domainpb/event_xxxx.proto.tmpl
var domainpbevent string

func generate(file *protomodels.File, plugin *protogen.Plugin) error {
	// 生成输出文件，文件名与源文件同名
	outputFile := plugin.NewGeneratedFile(
		fmt.Sprintf("%s.proto", file.GeneratedFilenamePrefix),
		file.GoImportPath,
	)

	// 处理文件中的消息
	for _, message := range file.Messages {
		// 检查消息名称是否以 "Event" 结尾
		if !strings.HasSuffix(message.Name, "Event") {
			continue
		}

		// 渲染模板并写入输出文件
		renderTemplate(outputFile, domainpbevent, message)
	}
	return nil
}

// renderTemplate 渲染模板并写入输出文件
func renderTemplate(outputFile *protogen.GeneratedFile, tmplContent string, data *protomodels.Message) {
	// 创建模板并加载子模板
	tmpl := plugins.NewTemplate("domain-event-tmpl")
	plugins.ParseProtoTemplate(tmpl, tmplContent)
	// 执行模板
	if err := tmpl.Execute(outputFile, data); err != nil {
		panic(fmt.Sprintf("failed to execute template: %v", err))
	}
}

func main() {
	if err := plugins.Plugin(generate); err != nil {
		panic(err)
	}
}

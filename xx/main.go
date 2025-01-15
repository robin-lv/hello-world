package main

import (
	_ "embed"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"protoc-plugins/protocplugins"
	"protoc-plugins/protocplugins/prototemplate"

	"strings"
)

//go:embed domainpb/event_xxxx.proto.tmpl
var domainpbevent string

func ProcessFile(file *prototemplate.File, plugin *protogen.Plugin) {
	// 生成输出文件，文件名与源文件同名
	outputFile := plugin.NewGeneratedFile(
		fmt.Sprintf("domain_%s.proto", strings.TrimSuffix(file.GeneratedFilenamePrefix, ".proto")),
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
}

// renderTemplate 渲染模板并写入输出文件
func renderTemplate(outputFile *protogen.GeneratedFile, tmplContent string, data *prototemplate.Message) {
	// 创建模板并加载子模板
	tmpl := prototemplate.New("domain-event-tmpl")

	// 加载主模板
	tmpl, err := tmpl.Parse(tmplContent)
	if err != nil {
		panic(fmt.Sprintf("failed to parse template: %v", err))
	}

	// 执行模板
	if err = tmpl.Execute(outputFile, data); err != nil {
		panic(fmt.Sprintf("failed to execute template: %v", err))
	}
}

func main() {
	protocplugins.Run(ProcessFile)
}

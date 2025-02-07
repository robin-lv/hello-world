package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"strings"
	"text/template"
)

const goTemplate = `// Code generated by protoc-gen-go-template. DO NOT EDIT.
// source: {{.Source}}

package {{.Package}}

// EventPublisher 是一个事件发布者，用于发布角色相关的事件。
type EventPublisher struct {
}

{{range .Messages}}
// Publish{{.Name}} 发布 {{.Name}} 事件。
func (p *EventPublisher) Publish{{.Name}}(evt *{{.Name}}) {
    // TODO: 实现事件发布逻辑
    // 例如：将事件发送到消息队列或调用其他服务
}
{{end}}
`

type TemplateData struct {
	Source   string
	Package  string
	Messages []MessageData
}

type MessageData struct {
	Name string
}

func main() {
	// 读取 protoc 的输入
	input, err := os.ReadFile("input.bin") // 替换为实际的输入文件
	if err != nil {
		panic(fmt.Sprintf("failed to read input: %v", err))
	}

	// 解析输入为 CodeGeneratorRequest
	var req pluginpb.CodeGeneratorRequest
	if err = proto.Unmarshal(input, &req); err != nil {
		panic(fmt.Sprintf("failed to parse input: %v", err))
	}

	// 初始化插件
	plugin, err := protogen.Options{}.New(&req)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize plugin: %v", err))
	}

	// 准备模板数据
	data := TemplateData{
		Source:  req.GetParameter(),
		Package: "role_domain_pb", // 替换为实际的包名
		Messages: []MessageData{
			{Name: "SetNameEvent"},
			// 添加更多消息
		},
	}

	// 解析模板
	tmpl, err := template.New("go").Parse(goTemplate)
	if err != nil {
		panic(fmt.Sprintf("failed to parse template: %v", err))
	}

	// 生成 Go 文件
	outputFile := "event_publisher.go" // 替换为实际的输出文件
	file, err := os.Create(outputFile)
	if err != nil {
		panic(fmt.Sprintf("failed to create output file: %v", err))
	}
	defer file.Close()

	// 执行模板
	if err = tmpl.Execute(file, data); err != nil {
		panic(fmt.Sprintf("failed to execute template: %v", err))
	}

	fmt.Println("Go file generated successfully:", outputFile)
}

//go:embed ../protoc-plugins/example/gen_doamin_event/event_xxxx.proto.tmpl
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
	tmpl := proto_plugins2.NewTemplate("domain-event-tmpl")
	proto_plugins2.ParseProtoTemplate(tmpl, tmplContent)
	// 执行模板
	if err := tmpl.Execute(outputFile, data); err != nil {
		panic(fmt.Sprintf("failed to execute template: %v", err))
	}
}

func main() {
	if err := proto_plugins.Plugin(generate); err != nil {
		panic(err)
	}
}

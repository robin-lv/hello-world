package protocplugins

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
	"protoc-plugins/protocplugins/prototemplate"
)

// genFunc 定义插件接口
type genFunc func(file *prototemplate.File, plugin *protogen.Plugin)

// Run 运行 Protobuf 插件
func Run(fn genFunc) {
	// 读取 protoc 的输入
	input, err := io.ReadAll(os.Stdin)
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

	// 处理所有 .proto 文件
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}
		fn(prototemplate.ProcessFile(file), plugin)
	}

	// 生成响应
	resp := plugin.Response()

	// 将响应序列化为二进制数据
	respData, err := proto.Marshal(resp)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal response: %v", err))
	}

	// 将二进制数据写入标准输出
	if _, err = os.Stdout.Write(respData); err != nil {
		panic(fmt.Sprintf("failed to write response: %v", err))
	}
}

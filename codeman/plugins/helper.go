package plugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"path/filepath"
	"protoc-plugins/plugins/protomodels"
	"strings"
)

type CommandHandle func(cmd *cobra.Command, args []string) error

type GenerateFunc func(file *protomodels.File, plugin *protogen.Plugin) error

func newGeneratorCommandFunc(req *pluginpb.CodeGeneratorRequest, p GenerateFunc) CommandHandle {
	return func(cmd *cobra.Command, args []string) error {
		plugin, err := protogen.Options{}.New(req)
		if err != nil {
			return errors.Wrap(err, "failed to initialize plugin")
		}

		// 处理所有 .proto 文件
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			if err = p(protomodels.ProcessFile(file), plugin); err != nil {
				return err
			}
		}

		resp := plugin.Response()

		// 将响应序列化为二进制数据
		respData, err := proto.Marshal(resp)
		if err != nil {
			return errors.Wrap(err, "failed to marshal response")
		}

		// 将二进制数据写入标准输出
		if _, err = os.Stdout.Write(respData); err != nil {
			return errors.Wrap(err, "failed to write response")
		}

		// 如果没有错误，返回 nil
		return nil
	}
}

// 获取插件名称
func getPluginName() string {
	executablePath := os.Args[0]
	executableName := filepath.Base(executablePath)
	if strings.HasPrefix(executableName, "protoc-gen-") {
		return strings.TrimPrefix(executableName, "protoc-gen-")
	}
	return executableName
}

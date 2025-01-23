package plugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
	"protoc-plugins/plugins/protomodels"
	"strings"
)

func NewGenerator() (g *Generator, err error) {
	g = &Generator{}
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		err = errors.Wrap(err, "failed to read input")
		return
	}

	if err = proto.Unmarshal(input, &g.input); err != nil {
		err = errors.Wrap(err, "failed to parse input")
		return
	}

	// 获取插件参数
	pluginParams := g.input.GetParameter()

	// 将参数解析为 os.Args 的形式
	args := strings.Fields(pluginParams)
	os.Args = append(os.Args[:1], args...)

	g.plugin, err = protogen.Options{}.New(&g.input)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize plugin")
		return
	}
	g.root = &cobra.Command{
		Use: getPluginName(),
	}
	return
}

type Generator struct {
	input  pluginpb.CodeGeneratorRequest
	plugin *protogen.Plugin
	root   *cobra.Command
}

func (g *Generator) AddCommands(cmd *cobra.Command, p process) {
	cmd.RunE = g.newCommandFunc(p)
	g.root.AddCommand(cmd)
}

// process 定义插件接口
type process func(file *protomodels.File, plugin *protogen.Plugin) error

func (g *Generator) newCommandFunc(p process) (run func(cmd *cobra.Command, args []string) error) {
	return func(cmd *cobra.Command, args []string) error {
		// 处理所有 .proto 文件
		for _, file := range g.plugin.Files {
			if !file.Generate {
				continue
			}
			if err := p(protomodels.ProcessFile(file), g.plugin); err != nil {
				return err
			}
		}

		// 生成响应
		resp := g.plugin.Response()

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

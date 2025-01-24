package plugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
	"strings"
)

func NewGenerator() (g *Generator, err error) {
	g = &Generator{}
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		err = errors.Wrap(err, "failed to read input")
		return
	}

	if err = proto.Unmarshal(input, &g.request); err != nil {
		err = errors.Wrap(err, "failed to parse request")
		return
	}

	return
}

type Generator struct {
	request pluginpb.CodeGeneratorRequest
}

func (g *Generator) Request() *pluginpb.CodeGeneratorRequest { return &g.request }

func (g *Generator) NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: getPluginName(),
	}
	// 获取插件参数
	pluginParams := g.request.GetParameter()
	args := strings.Fields(pluginParams)
	cmd.SetArgs(args)
	return cmd
}
func (g *Generator) WarpCommand(cmd *cobra.Command) *cobra.Command {

}

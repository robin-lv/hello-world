package cobrax

import (
	_ "embed"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	cobracompletefig "github.com/withfig/autocomplete-tools/integrations/cobra"
	"os"
	"protoc-plugins/internal/gen"
	"runtime"
	"text/template"
)

const (
	codeFailure = 1
)

var (
	//go:embed usage.tpl
	usageTpl string
)

// Execute executes the given command
func Execute(rootCmd *Command, args []string) {
	rootCmd.MustInit()
	rootCmd.Command.AddCommand(cobracompletefig.CreateCompletionSpecCommand())
	args = supportGoStdFlag(rootCmd, args)
	os.Args = append(os.Args[1:], args...)
	//rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.Red.Render(err.Error()))
		os.Exit(codeFailure)
	}
}

func MakeRootCommand(name, version string) *Command {
	rootCmd := NewCommand(name)
	cobra.AddTemplateFuncs(template.FuncMap{
		"blue":    blue,
		"green":   green,
		"rpadx":   rpadx,
		"rainbow": rainbow,
	})
	rootCmd.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)

	rootCmd.SetUsageTemplate(usageTpl)
	rootCmd.AddCommand(gen.Cmd)
	return rootCmd
}

package cmd

import (
	"github.com/spf13/cobra"
)

func GenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate code from .proto files",
	}
}

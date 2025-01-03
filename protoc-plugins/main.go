package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	outputDir  string
	customOpt  string
	protoFiles []string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "protoc-gen-custom",
		Short: "A custom protoc plugin",
		Long:  "A custom protoc plugin to generate code based on .proto files",
		Run: func(cmd *cobra.Command, args []string) {
			// Read the CodeGeneratorRequest from stdin
			input, err := io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatalf("Failed to read input: %v", err)
			}

			// Parse the CodeGeneratorRequest
			request := &pluginpb.CodeGeneratorRequest{}
			if err := proto.Unmarshal(input, request); err != nil {
				log.Fatalf("Failed to parse CodeGeneratorRequest: %v", err)
			}

			// Generate code
			response := generateCode(request)

			// Write the CodeGeneratorResponse to stdout
			output, err := proto.Marshal(response)
			if err != nil {
				log.Fatalf("Failed to marshal CodeGeneratorResponse: %v", err)
			}
			os.Stdout.Write(output)
		},
	}

	// Add flags
	rootCmd.PersistentFlags().StringVar(&outputDir, "output_dir", ".", "Output directory for generated files")
	rootCmd.PersistentFlags().StringVar(&customOpt, "custom_opt", "", "Custom option for the plugin")
	rootCmd.PersistentFlags().StringSliceVar(&protoFiles, "proto_files", []string{}, "List of .proto files")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func generateCode(request *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	response := &pluginpb.CodeGeneratorResponse{}

	// Process each .proto file
	for _, file := range request.GetProtoFile() {
		// Get the file name without extension
		fileName := file.GetName()
		outputFileName := fmt.Sprintf("%s/%s_custom_output.txt", outputDir, fileName[:len(fileName)-len(".proto")])

		// Generate content
		content := fmt.Sprintf("Generated code for %s with custom option: %s\n", fileName, customOpt)

		// Add the generated file to the response
		response.File = append(response.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(outputFileName),
			Content: proto.String(content),
		})
	}

	return response
}

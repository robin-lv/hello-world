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

var rootCmd = &cobra.Command{
	Use:   "protoc-gen-nine",
	Short: "A protoc plugin named 'nine'",
	Long:  "A custom protoc plugin named 'nine' to generate code based on .proto files",
}

func main() {

	// Add generate subcommand
	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate code from .proto files",
		Long:  "Generate code from .proto files using the protoc-gen-nine plugin",
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

	// Add generate subcommand to the root command
	rootCmd.AddCommand(generateCmd)

	// Execute the root command
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
		outputFileName := fmt.Sprintf("%s/%s_nine_output.txt", outputDir, fileName[:len(fileName)-len(".proto")])

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

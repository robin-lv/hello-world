package plugins

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取插件名称
func getPluginName() string {
	executablePath := os.Args[0]
	executableName := filepath.Base(executablePath)
	if strings.HasPrefix(executableName, "protoc-gen-") {
		return strings.TrimPrefix(executableName, "protoc-gen-")
	}
	return executableName
}

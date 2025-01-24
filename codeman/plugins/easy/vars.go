package easy

import (
	_ "embed"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func ProtocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

type CommentSet struct {
	Head []protogen.Comments
	Body string
	Foot string
}

func MakeCommentSet(p protogen.CommentSet) CommentSet {
	return CommentSet{
		Foot: strings.TrimSpace(p.Trailing.String()),
		Body: strings.TrimSpace(p.Leading.String()),
		Head: p.LeadingDetached,
	}
}

type Field struct {
	Name    string
	Num     int
	Comment CommentSet
}

type FileHeadInfo struct {
	Plugin    string
	PluginVer string
	ProtocVer string
	Source    string
	Syntax    string
	GoPackage string
}

func MakeFileHeadInfo(pluginName, pluginVer string, gen *protogen.Plugin, file *protogen.File) *FileHeadInfo {
	return &FileHeadInfo{
		Plugin:    pluginName,
		PluginVer: pluginVer,
		ProtocVer: ProtocVersion(gen),
		Source:    file.Desc.Path(),
		Syntax:    file.Proto.GetSyntax(),
		GoPackage: file.Proto.Options.GetGoPackage(),
	}
}

package prototemplate

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// File 是处理后的文件结构
type File struct {
	*protogen.File
	Messages []*Message // 处理后的消息
	Enums    []*Enum    // 处理后的枚举
	Services []*Service // 处理后的服务
}

// Message 是处理后的消息结构
type Message struct {
	*protogen.Message
	Name           string     // 存储原始的 Protobuf 消息名称
	Fields         []*Field   // 处理后的字段
	NestedMessages []*Message // 处理后的嵌套消息
	Enums          []*Enum    // 处理后的枚举
	Oneofs         []*Oneof   // 处理后的 oneof
}

// Field 是处理后的字段结构
type Field struct {
	*protogen.Field
	Name string // 存储原始的 Protobuf 字段名称
}

// Enum 是处理后的枚举结构
type Enum struct {
	*protogen.Enum
	Name   string       // 存储原始的 Protobuf 枚举名称
	Values []*EnumValue // 处理后的枚举值
}

// EnumValue 是处理后的枚举值结构
type EnumValue struct {
	*protogen.EnumValue
	Name string // 存储原始的 Protobuf 枚举值名称
}

// Oneof 是处理后的 oneof 结构
type Oneof struct {
	*protogen.Oneof
	Name string // 存储原始的 Protobuf oneof 名称
}

// Service 是处理后的服务结构
type Service struct {
	*protogen.Service
	Name    string    // 存储原始的 Protobuf 服务名称
	Methods []*Method // 处理后的方法
}

// Method 是处理后的方法结构
type Method struct {
	*protogen.Method
	Name string // 存储原始的 Protobuf 方法名称
}

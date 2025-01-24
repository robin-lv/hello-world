package models

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// ProcessFile 将 protogen.File 转换为 File
func ProcessFile(file *protogen.File) *File {
	processed := &File{
		File:     file,
		Messages: make([]*Message, 0, len(file.Messages)),
		Enums:    make([]*Enum, 0, len(file.Enums)),
		Services: make([]*Service, 0, len(file.Services)),
	}

	// 处理消息
	for _, msg := range file.Messages {
		processed.Messages = append(processed.Messages, processMessage(msg))
	}

	// 处理枚举
	for _, enum := range file.Enums {
		processed.Enums = append(processed.Enums, processEnum(enum))
	}

	// 处理服务
	for _, service := range file.Services {
		processed.Services = append(processed.Services, processService(service))
	}

	return processed
}

// processMessage 将 protogen.Message 转换为 Message
func processMessage(msg *protogen.Message) *Message {
	processed := &Message{
		Message:        msg,
		Name:           string(msg.Desc.Name()), // 获取原始消息名称
		Fields:         make([]*Field, 0, len(msg.Fields)),
		NestedMessages: make([]*Message, 0, len(msg.Messages)),
		Enums:          make([]*Enum, 0, len(msg.Enums)),
		Oneofs:         make([]*Oneof, 0, len(msg.Oneofs)),
	}

	// 处理字段
	for _, field := range msg.Fields {
		processed.Fields = append(processed.Fields, &Field{
			Field: field,
			Name:  string(field.Desc.Name()),
		})
	}

	// 处理嵌套消息
	for _, nestedMsg := range msg.Messages {
		processed.NestedMessages = append(processed.NestedMessages, processMessage(nestedMsg))
	}

	// 处理枚举
	for _, enum := range msg.Enums {
		processed.Enums = append(processed.Enums, processEnum(enum))
	}

	// 处理 oneof
	for _, oneof := range msg.Oneofs {
		processed.Oneofs = append(processed.Oneofs, &Oneof{
			Oneof: oneof,
			Name:  string(oneof.Desc.Name()),
		})
	}

	return processed
}

// processEnum 将 protogen.Enum 转换为 Enum
func processEnum(enum *protogen.Enum) *Enum {
	processed := &Enum{
		Enum:   enum,
		Name:   string(enum.Desc.Name()),
		Values: make([]*EnumValue, 0, len(enum.Values)),
	}

	// 处理枚举值
	for _, value := range enum.Values {
		processed.Values = append(processed.Values, &EnumValue{
			EnumValue: value,
			Name:      string(value.Desc.Name()),
		})
	}

	return processed
}

// processService 将 protogen.Service 转换为 Service
func processService(service *protogen.Service) *Service {
	processed := &Service{
		Service: service,
		Name:    string(service.Desc.Name()),
		Methods: make([]*Method, 0, len(service.Methods)),
	}

	// 处理方法
	for _, method := range service.Methods {
		processed.Methods = append(processed.Methods, &Method{
			Method: method,
			Name:   string(method.Desc.Name()),
		})
	}

	return processed
}

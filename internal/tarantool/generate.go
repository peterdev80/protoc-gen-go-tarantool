package tarantool

import "google.golang.org/protobuf/compiler/protogen"

// MM описывает имя соответствие типа.
type MM map[string][]string

// Last вернуть последний элемент.
func (m MM) Last(key string) string {
	lst, ok := m[key]
	if ok {
		return lst[len(lst)-1]
	}

	return ""
}

// MapCmd  proto-golang`s возможные типы которые могут прийти из interface,
// замыкающий тип slice соответствует типу поля.
var MapCmd = MM{ //nolint:gochecknoglobals
	"int32":   {"int8", "int16", "int32"},
	"int64":   {"int8", "int16", "int32", "int64"},
	"float":   {"float32"},
	"double":  {"float32", "float64"},
	"uint32":  {"uint8", "uint16", "uint32"},
	"uint64":  {"uint8", "uint16", "uint32", "uint64"},
	"sint32":  {"int8", "int16", "int32"},
	"sint64":  {"int8", "int16", "int32", "int64"},
	"fixed32": {"uint8", "uint16", "uint32"},
	// TO-DO describe all scalar type
	"string": {"string"},
	"bool":   {"bool"},
	"bytes":  {"[]byte"},
}

// typeCmd вернуть в зависимости от типа [encoder,decoder,go type].
func typeCmd(tp string) (string, string, string) { //nolint:cyclop
	switch tp {
	case "string":
		return "EncodeString", "DecodeString", "string"
	case "uint32":
		return "EncodeUint32", "DecodeUint32", "uint32"

	case "double":
		return "EncodeFloat64", "DecodeFloat64", "float64"
	case "float":
		return "EncodeFloat32", "DecodeFloat32", "float32"
	case "int32":
		return "EncodeInt32", "DecodeInt32", "int32"
	case "int64":
		return "EncodeInt64", "DecodeInt64", "int64"
	case "uint64":
		return "EncodeUint64", "DecodeUint64", "uint64"
	case "bool":
		return "EncodeBool", "DecodeBool", "bool"
	case "byte":
		return "EncodeBytes", "DecodeBytes", "[]byte"
	default:
		return "Encode", "Decode", ""

	}
}

// Вернуть количество полей с учетом, что oneof одно поле
func lenGen(msg *protogen.Message) int {
	// исключаем из len one off
	ln := len(msg.Fields)
	for _, of := range msg.Oneofs {
		ln -= len(of.Fields)
		// одно поле выбрано
		ln++
	}

	return ln
}

// kind Вернуть тип поля или тип mesage.
func kind(f *protogen.Field) string {
	k := f.Desc.Kind().String()
	if f.Message != nil {
		k = f.Message.GoIdent.GoName
	}

	return k
}

func GeneratorMsgPK(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + ".msgpk.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-ascii. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P(`import (
    "fmt"
	"github.com/vmihailenco/msgpack/v5"
	)`)
	g.P("//", file.Messages[0])

	for _, msg := range file.Messages {
		g.P("func (x *", msg.GoIdent, ")EncodeMsgpack(enc *msgpack.Encoder) error {")

		encoder(g, msg)
		g.P("return nil")
		g.P("}")

		g.P("func (x *", msg.GoIdent, ")DecodeMsgpack(dec *msgpack.Decoder) error {")
		decode(g, msg)
		g.P("return nil")
		g.P("}")
	}

	return g
}

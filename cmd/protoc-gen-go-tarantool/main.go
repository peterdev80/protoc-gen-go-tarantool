package main

import (
	"github.com/peterdev80/protoc-gen-go-tarantool/internal/tarantool"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			tarantool.GeneratorMsgPK(gen, f)
		}
		return nil
	})
}
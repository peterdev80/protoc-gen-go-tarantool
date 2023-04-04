# protoc-gen-go-tarantool
protoc генератор сообщений для tarantool (msgpack)

Plugin для protoc для генерации реализации интерфейсов Encoder Decoder пакета messagepack (github.com/vmihailenco/msgpack/v5) для обработки в tarantool.

На машине необходимо установить 
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
```
(https://protobuf.dev/reference/go/go-generated/)


установить плагин 
```
go install github.com/peterdev80/protoc-gen-go-tarantool/cmd/protoc-gen-go-tarantool@latest
```

для генерации структуры для tarantool 
```
	protoc \
    	  --proto_path=$(PROTO_PATCH)\
           --go_out=. \
           --go-tarantool_out=. \
         $(PROTO_PATCH)/a1/*.proto
```

Демонстрация работы в [example](./example/README.md)
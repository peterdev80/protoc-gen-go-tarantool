PROTO = proto

proto:
	protolint -fix $(PROTO)
	protoc \
	--proto_path=$(PROTO) \
	--go_out=. \
	--go-tarantool_out=. \
	--doc_opt=markdown,API.md,source_relative \
	--doc_out=. \
	$(PROTO)/*.proto

lint:
	golangci-lint run --fix ./...


run:
	go run -tags go_tarantool_ssl_disable,go_tarantool_msgpack_v5 main.go

.PHONY: lint proto run
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

.PHONY: lint proto
GO = GO111MODULE=on go
PROTOC = protoc --plugin=./bin/protoc-gen-go

.PHONY: all clean golang

all: clean golang

clean:
	rm -rf go/datasource/*.pb.go

protoc-gen-go: tools/go.mod
	@cd tools; \
	$(GO) build -o ../bin/protoc-gen-go github.com/golang/protobuf/protoc-gen-go

golang: protoc-gen-go
	${PROTOC} -I . datasource.proto --go_out=plugins=grpc:go/datasource/.
	${PROTOC} -I . renderer.proto --go_out=plugins=grpc:go/renderer/.

all: clean golang

clean:
	rm -rf go/models/*.pb.go

golang:
	protoc -I . datasource.proto --go_out=plugins=grpc:go/models/.
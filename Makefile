all: clean golang

clean:
	rm -rf go/datasource/*.pb.go
	rm -rf go/renderer/*.pb.go

golang:
	protoc -I . datasource.proto --go_out=plugins=grpc:go/datasource/.
	protoc -I . renderer.proto --go_out=plugins=grpc:go/renderer/.

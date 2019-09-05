all: clean golang

clean:
	rm -rf go/datasource/*.pb.go
	rm -rf go/renderer/*.pb.go

golang:
	protoc -I proto datasource.proto --go_out=plugins=grpc:go/datasource/.
	protoc -I proto renderer.proto --go_out=plugins=grpc:go/renderer/.

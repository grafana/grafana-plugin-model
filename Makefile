all: clean golang

clean:
	rm -rf go/datasource/*.pb.go

golang:
	./genproto.sh
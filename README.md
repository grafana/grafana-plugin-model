# Models for Grafana backend plugins

## Install tooling

```bash
cd /tmp
curl -OL https://github.com/google/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip
unzip protoc-3.9.1-linux-x86_64.zip -d protoc3
sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/
sudo chown root /usr/local/bin/protoc
sudo chown -R root /usr/local/include/google
```

```bash
go get -d -u github.com/golang/protobuf/protoc-gen-go
cd $GOPATH/src/github.com/google/protobuf/protoc-gen-go
git checkout 927b65914520a8b7d44f5c9057611cfec6b2e2d0
go install github.com/golang/protobuf/protoc-gen-go
```

## Build

Generate go code using the protocol buffer compiler, protoc, on .proto files.

```bash
make
````
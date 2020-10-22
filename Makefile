# vi: ft=make
.PHONY: proto test

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . lile_test.proto --lile-server_out=. --go_out=plugins=grpc,paths=source_relative:.

test: proto
	go test -p 1 -v ./...

.PHONY: protoc-setup
protoc-setup:
	docker build -t memo-protc ../build
	docker run --rm -v `pwd`:/go/src/github.com/hikaru7719/memo/server memo-protc

.PHONY: protoc
protoc: protoc-grpc

.PHONY: protoc-grpc
protoc-grpc:
	protoc -I proto/ --go_out=plugins=grpc:proto/ proto/*.proto

.PHONY: cp-proto
cp-proto:
	cp ../proto/*.proto proto/


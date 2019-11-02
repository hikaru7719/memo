.PHONY: protoc-setup
protoc-setup:
	docker build -t memo-protc ../build
	docker run --rm -v `pwd`:/go/src/github.com/hikaru7719/memo/server memo-protc

.PHONY: protoc
protoc: protoc-grpc protoc-grpc-gateway

.PHONY: protoc-grpc
protoc-grpc:
	protoc -I proto/ \
	-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:proto/ \
	proto/*.proto

.PHONY: protoc-grpc-gateway
protoc-grpc-gateway:
	protoc -I proto/ \
	-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:proto/ \
	proto/*.proto

.PHONY: cp-proto
cp-proto:
	cp ../proto/*.proto proto/


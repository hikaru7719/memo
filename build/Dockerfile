FROM golang:1.13
SHELL ["/bin/bash", "-c"]
RUN apt-get update && apt-get install -y --no-install-recommends \
    autoconf \
    automake \
    libtool

ENV PROTOBUF_VERSION 3.10.1
RUN wget -O protobuf.tar.gz "https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/protobuf-cpp-${PROTOBUF_VERSION}.tar.gz"
RUN tar xzf protobuf.tar.gz && \
    cd "protobuf-${PROTOBUF_VERSION}" && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install &&\
    ldconfig

ENV GO_PROTOBUF_VERSION 1.3.2
RUN go get -u github.com/golang/protobuf/protoc-gen-go && \
    git -C /go/src/github.com/golang/protobuf checkout "v${GO_PROTOBUF_VERSION}" && \
    go install github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/github.com/hikaru7719/memo/server

CMD ["make", "protoc"]

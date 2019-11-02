FROM golang:1.13 as builder
COPY . /grpc
WORKDIR /grpc
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/grpc

FROM alpine:latest
WORKDIR /app
COPY --from=builder /grpc .
CMD ["./grpc"]

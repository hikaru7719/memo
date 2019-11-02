FROM golang:1.13 as builder
COPY . /gateway
WORKDIR /gateway
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/gateway

FROM alpine:latest
WORKDIR /app
COPY --from=builder /gateway .
CMD ["./gateway"]

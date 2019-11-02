package main

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/hikaru7719/memo/server/proto"
	"github.com/hikaru7719/memo/server/usecase"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHealthcheckServiceServer(grpcServer, &usecase.HealthcheckServiceImpl{})
	grpcServer.Serve(lis)
}

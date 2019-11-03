package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/hikaru7719/memo/server/controller"
	pb "github.com/hikaru7719/memo/server/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHealthcheckServiceServer(grpcServer, &controller.HealthcheckServiceImpl{})
	grpcServer.Serve(lis)
}

package usecase

import (
	"context"
	"fmt"

	pb "github.com/hikaru7719/memo/server/proto"
)

// HealthcheckServiceImpl is for healthcheck
type HealthcheckServiceImpl struct{}

// SayHello is for healthcheck
func (h *HealthcheckServiceImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println(req.Greeting)
	return &pb.HelloResponse{Reply: "OK"}, nil
}

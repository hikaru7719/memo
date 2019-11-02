package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/hikaru7719/memo/server/proto"
)

// gRPCServerEndpoint is port of grpc server
// This Var should move to configulation files
const gRPCServerEndpoint = ":8080"

func run() error {
	ctx := context.Background()
	ctx, canccel := context.WithCancel(ctx)
	defer canccel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHealthcheckServiceHandlerFromEndpoint(ctx, mux, gRPCServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

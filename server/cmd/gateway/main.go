package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/hikaru7719/memo/server/proto"
)

// gRPCServerEndpoint is port of grpc server
// This Var should move to configulation files
const gRPCServerEndpoint = "grpc:8080"

func run() error {
	ctx := context.Background()
	ctx, canccel := context.WithCancel(ctx)
	defer canccel()

	mux := runtime.NewServeMux()
	newMux := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{"content-type", "application/json"}),
	)(mux)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHealthcheckServiceHandlerFromEndpoint(ctx, mux, gRPCServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", newMux)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

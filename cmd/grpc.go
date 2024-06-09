package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_sentry "github.com/johnbellone/grpc-middleware-sentry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func initRPC() *grpc.Server {
	s := grpc.NewServer(
		grpc.Creds(nil),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_sentry.UnaryServerInterceptor(),
		)),
	)
	reflection.Register(s)
	return s
}

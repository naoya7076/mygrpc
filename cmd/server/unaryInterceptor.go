package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func myUnaryServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] Unary interceptor 1: ", info.FullMethod)
	res, err := handler(ctx, req)
	log.Println("[post] Unary interceptor 1: ")
	return res, err
}

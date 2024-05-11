package grpc

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

func TimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		childCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return handler(childCtx, req)
	}
}

package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func AuthorizationServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	h, err := handler(ctx, req)
	grpclog.Infof("Some authorization:%s", info.FullMethod)
	return h, err
}

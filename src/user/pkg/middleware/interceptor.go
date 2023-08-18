package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
	"user.com/internal/auth"
)

type Interceptor struct {
	auth auth.AuthServiceClient
}

func NewInterceptor(auth auth.AuthServiceClient) *Interceptor {
	return &Interceptor{
		auth: auth,
	}
}

func (i *Interceptor) ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	if err := i.authorize(ctx); err != nil {
		return nil, err
	}

	h, err := handler(ctx, req)
	grpclog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)
	return h, err
}

func (i *Interceptor) authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["api-token"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]
	err := i.auth.AuthServiceToken(token)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	return nil
}

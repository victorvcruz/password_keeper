package utils

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"strings"
)

func BearerToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("without metadata")
	}

	token := md.Get("authorization")
	if token == nil {
		return "", errors.New("without token")
	}

	return strings.ReplaceAll(token[0], "Bearer ", ""), nil
}

func GetMetadata(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	value := md.Get(key)
	if value == nil {
		return ""
	}

	return md.Get(key)[0]
}

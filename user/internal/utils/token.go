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

	token := md.Get("authorization")[0]
	if token == "" {
		return "", errors.New("without token")
	}

	return strings.ReplaceAll(token, "Bearer ", ""), nil
}


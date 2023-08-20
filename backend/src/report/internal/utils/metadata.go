package utils

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func GetMetadataByKey(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("without metadata")
	}

	value := md.Get(key)[0]
	if value == "" {
		return "", errors.New(fmt.Sprintf("without %s", key))
	}

	return value, nil
}

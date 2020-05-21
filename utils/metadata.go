package utils

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// GetUserIDMetadata retrieves userID from context metadata, returns grpc errors
func GetUserIDMetadata(ctx context.Context) (int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Errorf(codes.DataLoss, "Failed to get userID metadata")
	}
	if userIDMap, ok := md["userid"]; ok {
		userID, err := strconv.ParseInt(userIDMap[0], 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}
	return 0, status.Errorf(codes.DataLoss, "Failed to get userID metadata")
}

// PassUserIDMetadata passes the given userID into context metadata, returns new context
func PassUserIDMetadata(ctx context.Context, userID string) context.Context {
	md := metadata.Pairs("userid", userID)
	return metadata.NewIncomingContext(ctx, md)
}

package keeper

import (
	"context"

	"voter/x/voter/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ShowPoll(ctx context.Context, req *types.QueryShowPollRequest) (*types.QueryShowPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryShowPollResponse{}, nil
}

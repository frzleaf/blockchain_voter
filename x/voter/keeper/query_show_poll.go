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

	poll, found := q.k.GetPoll(ctx, req.PollId)
	if !found {
		return nil, status.Error(codes.NotFound, "poll not found")
	}

	return &types.QueryShowPollResponse{
		Creator: poll.Creator,
		Id:      poll.Id,
		Title:   poll.Title,
		Options: poll.Options,
	}, nil
}

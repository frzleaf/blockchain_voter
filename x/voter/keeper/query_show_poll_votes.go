package keeper

import (
	"context"

	"voter/x/voter/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ShowPollVotes(ctx context.Context, req *types.QueryShowPollVotesRequest) (*types.QueryShowPollVotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(q.k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.VoteKey))

	var votes []*types.Vote
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var vote types.Vote
		if err := q.k.cdc.Unmarshal(value, &vote); err != nil {
			return err
		} else if vote.PollID == req.PollId {
			votes = append(votes, &vote)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryShowPollVotesResponse{
		Votes:      votes,
		Pagination: pageRes,
	}, nil
}

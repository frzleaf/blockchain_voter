package keeper

import (
	"context"
	"encoding/binary"

	"voter/x/voter/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CastVote(ctx context.Context, vote types.Vote) error {
	_, found := k.GetPoll(ctx, vote.PollID)
	if !found {
		return errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "poll not found")
	}

	votes := k.GetAllVote(ctx)
	for _, existingVote := range votes {
		if existingVote.Creator == vote.Creator && existingVote.PollID == vote.PollID {
			return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "already voted on this poll")
		}
	}

	count := k.GetVoteCount(ctx)
	vote.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.VoteKey))
	appendedValue := k.cdc.MustMarshal(&vote)
	store.Set(ConvertUint64ToBytes(vote.Id), appendedValue)
	k.SetVoteCount(ctx, count+1)

	return nil
}

func (k Keeper) GetVoteCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := []byte(types.VoteCountKey)
	if bz := store.Get(byteKey); bz != nil {
		return binary.BigEndian.Uint64(bz)
	}
	return 0
}

func (k Keeper) SetVoteCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set([]byte(types.VoteCountKey), bz)
}

func (k Keeper) GetAllVote(ctx context.Context) (res []types.Vote) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.VoteKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var vote types.Vote
		k.cdc.MustUnmarshal(iterator.Value(), &vote)
		res = append(res, vote)
	}
	return
}

func (k Keeper) GetAllVotes(ctx context.Context, pollID uint64) (res []types.Vote) {
	allVotes := k.GetAllVote(ctx)
	for _, vote := range allVotes {
		if vote.PollID == pollID {
			res = append(res, vote)
		}
	}
	return
}

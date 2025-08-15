package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"voter/x/voter/types"
)

func (k Keeper) AppendPoll(ctx context.Context, poll types.Poll) uint64 {
	count := k.GetPollCount(ctx)
	poll.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PollKey))
	appendedValue := k.cdc.MustMarshal(&poll)
	store.Set(ConvertUint64ToBytes(poll.Id), appendedValue)
	k.SetPollCount(ctx, poll.Id+1)
	return count
}

func (k Keeper) GetPollCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := []byte(types.PollCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}
func ConvertUint64ToBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetPollCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := []byte(types.PollCountKey)
	store.Set(byteKey, ConvertUint64ToBytes(count))
}

func (k Keeper) GetPoll(ctx context.Context, id uint64) (res types.Poll, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PollKey))
	if b := store.Get(ConvertUint64ToBytes(id)); b != nil {
		k.cdc.MustUnmarshal(b, &res)
		found = true
	}
	return
}

func (k Keeper) GetAllPolls(ctx context.Context) (list []types.Poll) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PollKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var poll types.Poll
		k.cdc.MustUnmarshal(iterator.Value(), &poll)
		list = append(list, poll)
	}
	return
}

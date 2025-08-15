package keeper

import (
	"context"

	"voter/x/voter/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) CastVote(ctx context.Context, msg *types.MsgCastVote) (*types.MsgCastVoteResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgCastVoteResponse{}, nil
}

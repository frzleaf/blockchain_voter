package keeper

import (
	"context"
	"strconv"

	"voter/x/voter/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CastVote(goCtx context.Context, msg *types.MsgCastVote) (*types.MsgCastVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pollId, err := strconv.ParseInt(msg.PollId, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid poll id")
	}

	vote := types.Vote{
		PollID:  uint64(pollId),
		Creator: msg.Creator,
		Option:  msg.Option,
	}

	if err := k.Keeper.CastVote(ctx, vote); err != nil {
		return nil, err
	}

	return &types.MsgCastVoteResponse{
		Id:     pollId,
		Option: vote.Option,
	}, nil
}

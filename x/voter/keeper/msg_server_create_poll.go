package keeper

import (
	"context"

	"voter/x/voter/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePoll(goCtx context.Context, msg *types.MsgCreatePoll) (*types.MsgCreatePollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleAcct := k.authKeeper.GetModuleAddress(types.ModuleName)
	if moduleAcct == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnknownAddress, "module account does not exist")
	}

	feeCoin, err := sdk.ParseCoinsNormalized("200token")
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee amount")
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator")
	}

	spendableCoins := k.bankKeeper.SpendableCoins(ctx, creator)
	if !spendableCoins.IsAllGTE(feeCoin) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFee, "not enough funds")
	}

	if len(msg.Options) < 2 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "poll must have at least two options")
	}

	poll := types.Poll{
		Creator: msg.Creator,
		Options: msg.Options,
		Title:   msg.Title,
	}

	id := k.AppendPoll(goCtx, poll)
	return &types.MsgCreatePollResponse{
		Id:    id,
		Title: poll.Title,
	}, nil
}

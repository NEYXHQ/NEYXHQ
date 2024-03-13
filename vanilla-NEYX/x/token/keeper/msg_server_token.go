package keeper

import (
	"context"
	"fmt"

	"github.com/Palo_Alt0/vanilla-NEYX/x/token/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var token = types.Token{
		Creator: msg.Creator,
		Name:    msg.Name,
		Supply:  msg.Supply,
	}

	id := k.AppendToken(
		ctx,
		token,
	)

	return &types.MsgCreateTokenResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var token = types.Token{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Supply:  msg.Supply,
	}

	// Checks that the element exists
	val, found := k.GetToken(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetToken(ctx, token)

	return &types.MsgUpdateTokenResponse{}, nil
}

func (k msgServer) DeleteToken(goCtx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetToken(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveToken(ctx, msg.Id)

	return &types.MsgDeleteTokenResponse{}, nil
}

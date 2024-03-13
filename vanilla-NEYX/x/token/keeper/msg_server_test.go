package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Palo_Alt0/vanilla-NEYX/testutil/keeper"
	"github.com/Palo_Alt0/vanilla-NEYX/x/token/keeper"
	"github.com/Palo_Alt0/vanilla-NEYX/x/token/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TokenKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

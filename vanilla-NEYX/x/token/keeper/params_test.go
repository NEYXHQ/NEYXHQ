package keeper_test

import (
	"testing"

	testkeeper "github.com/Palo_Alt0/vanilla-NEYX/testutil/keeper"
	"github.com/Palo_Alt0/vanilla-NEYX/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TokenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

package token

import (
	"github.com/Palo_Alt0/vanilla-NEYX/x/token/keeper"
	"github.com/Palo_Alt0/vanilla-NEYX/x/token/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the token
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, elem)
	}

	// Set token count
	k.SetTokenCount(ctx, genState.TokenCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TokenList = k.GetAllToken(ctx)
	genesis.TokenCount = k.GetTokenCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

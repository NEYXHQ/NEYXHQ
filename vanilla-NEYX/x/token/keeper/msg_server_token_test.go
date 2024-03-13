package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/Palo_Alt0/vanilla-NEYX/x/token/types"
)

func TestTokenMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateToken(ctx, &types.MsgCreateToken{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTokenMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateToken
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateToken{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateToken{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateToken{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateToken(ctx, &types.MsgCreateToken{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateToken(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTokenMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteToken
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteToken{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteToken{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteToken{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateToken(ctx, &types.MsgCreateToken{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteToken(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

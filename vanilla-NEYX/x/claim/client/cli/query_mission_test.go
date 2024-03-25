package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/client/cli"
	"github.com/ignite/modules/x/claim/types"
)

func (suite *QueryTestSuite) TestShowMission() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ClaimState.Missions

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		name string
		id   string
		args []string
		err  error
		obj  types.Mission
	}{
		{
			name: "should allow get",
			id:   fmt.Sprintf("%d", objs[0].MissionID),
			args: common,
			obj:  objs[0],
		},
		{
			name: "should return not found",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		suite.T().Run(tc.name, func(t *testing.T) {
			require.NoError(t, suite.Network.WaitForNextBlock())

			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMission(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
				return
			}

			require.NoError(t, err)
			var resp types.QueryGetMissionResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NotNil(t, resp.Mission)
			require.Equal(t,
				nullify.Fill(&tc.obj),
				nullify.Fill(&resp.Mission),
			)
		})
	}
}

func (suite *QueryTestSuite) TestListMission() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ClaimState.Missions

	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	suite.T().Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMission(), args)
			require.NoError(t, err)
			var resp types.QueryAllMissionResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Mission), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Mission),
			)
		}
	})
	suite.T().Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMission(), args)
			require.NoError(t, err)
			var resp types.QueryAllMissionResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Mission), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Mission),
			)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMission(), args)
		require.NoError(t, err)
		var resp types.QueryAllMissionResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Mission),
		)
	})
}

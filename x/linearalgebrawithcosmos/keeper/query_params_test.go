package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "linear-algebra-with-cosmos/testutil/keeper"
	"linear-algebra-with-cosmos/x/linearalgebrawithcosmos/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.LinearalgebrawithcosmosKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

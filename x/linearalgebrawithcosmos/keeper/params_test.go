package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "linear-algebra-with-cosmos/testutil/keeper"
	"linear-algebra-with-cosmos/x/linearalgebrawithcosmos/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LinearalgebrawithcosmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

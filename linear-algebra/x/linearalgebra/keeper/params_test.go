package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "linear-algebra/testutil/keeper"
	"linear-algebra/x/linearalgebra/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LinearalgebraKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

package linearalgebra_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "linear-algebra/testutil/keeper"
	"linear-algebra/testutil/nullify"
	"linear-algebra/x/linearalgebra"
	"linear-algebra/x/linearalgebra/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LinearalgebraKeeper(t)
	linearalgebra.InitGenesis(ctx, *k, genesisState)
	got := linearalgebra.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

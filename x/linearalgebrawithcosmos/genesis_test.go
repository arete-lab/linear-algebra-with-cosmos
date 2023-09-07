package linearalgebrawithcosmos_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "linear-algebra-with-cosmos/testutil/keeper"
	"linear-algebra-with-cosmos/testutil/nullify"
	"linear-algebra-with-cosmos/x/linearalgebrawithcosmos"
	"linear-algebra-with-cosmos/x/linearalgebrawithcosmos/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LinearalgebrawithcosmosKeeper(t)
	linearalgebrawithcosmos.InitGenesis(ctx, *k, genesisState)
	got := linearalgebrawithcosmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

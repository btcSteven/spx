package spx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "spx/testutil/keeper"
	"spx/testutil/nullify"
	"spx/x/spx"
	"spx/x/spx/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SpxKeeper(t)
	spx.InitGenesis(ctx, *k, genesisState)
	got := spx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

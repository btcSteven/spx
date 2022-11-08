package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "spx/testutil/keeper"
	"spx/x/spx/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SpxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

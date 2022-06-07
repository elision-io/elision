package keeper_test

import (
	"testing"

	testkeeper "github.com/elision-io/elision/testutil/keeper"
	"github.com/elision-io/elision/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

package onix_test

import (
	"testing"

	keepertest "github.com/OnixChain/onix/testutil/keeper"
	"github.com/OnixChain/onix/x/onix"
	"github.com/OnixChain/onix/x/onix/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OnixKeeper(t)
	onix.InitGenesis(ctx, *k, genesisState)
	got := onix.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}

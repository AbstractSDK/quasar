package qbank_test

import (
	"testing"

	keepertest "github.com/abag/quasarnode/testutil/keeper"
	"github.com/abag/quasarnode/testutil/nullify"
	"github.com/abag/quasarnode/x/qbank"
	"github.com/abag/quasarnode/x/qbank/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # genesis/test/state
	}

	ctx, keeper := keepertest.NewTestSetup(t).GetQbankKeeper()
	qbank.InitGenesis(ctx, keeper, genesisState)
	got := qbank.ExportGenesis(ctx, keeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Params, got.Params)

	// this line is used by starport scaffolding # genesis/test/assert
}

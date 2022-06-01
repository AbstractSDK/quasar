package keeper

import (
	"context"
	"errors"
	"testing"

	"github.com/abag/quasarnode/x/intergamm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var scenarios map[string]func(sdk.Context, *Keeper) func(*testing.T)

func init() {
	scenarios = make(map[string]func(sdk.Context, *Keeper) func(*testing.T))
}

func (ms msgServer) TestScenario(goCtx context.Context, msg *types.MsgTestScenario) (*types.MsgTestScenarioResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.Logger().Info("Running test scenario", "scenario", msg.Scenario)

	f, ok := scenarios[msg.Scenario]

	if !ok {
		return nil, errors.New("unknown test scenario")
	}

	return runTest(msg.Scenario, f(ctx, ms.k)), nil
}
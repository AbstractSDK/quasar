package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/quasarlabs/quasarnode/x/qbank/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TotalWithdraw fetch the total amount of tokens withdraw so far by the given user.
func (k Keeper) TotalWithdraw(goCtx context.Context, req *types.QueryTotalWithdrawRequest) (*types.QueryTotalWithdrawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, _ := k.GetTotalWithdrawAmt(ctx, req.UserAcc, req.VaultID)

	return &types.QueryTotalWithdrawResponse{Coins: val}, nil
}

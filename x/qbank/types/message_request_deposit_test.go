package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/quasarlabs/quasarnode/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRequestDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    MsgRequestDeposit
		err    error
		errMsg string
	}{
		{
			name: "invalid address",
			msg: MsgRequestDeposit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid vault id",
			msg: MsgRequestDeposit{
				Creator: sample.AccAddressStr(),
				VaultID: "xyz",
			},
			errMsg: "invalid vault",
		}, {
			name: "valid address",
			msg: MsgRequestDeposit{
				Creator:      sample.AccAddressStr(),
				VaultID:      "orion",
				Coin:         sdk.NewCoin("abc", sdk.NewInt(100)),
				LockupPeriod: 1,
			},
		},
		{
			name: "valid reserved length",
			msg: MsgRequestDeposit{
				Creator:      sample.AccAddressStr(),
				VaultID:      "orion",
				Coin:         sdk.NewCoin("abc", sdk.NewInt(100)),
				LockupPeriod: 1,
				Reserved:     []string{"abc"},
			},
			errMsg: "invalid reserved field length",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if err != nil {
				if tt.err != nil {
					require.ErrorIs(t, err, tt.err)
					return
				}
				if tt.errMsg != "" {
					require.Equal(t, err.Error(), tt.errMsg)
					return
				}
			}
			require.NoError(t, err)
		})
	}
}

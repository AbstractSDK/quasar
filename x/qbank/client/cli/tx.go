package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/quasarlabs/quasarnode/x/qbank/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRequestDeposit())
	cmd.AddCommand(CmdRequestWithdraw())
	cmd.AddCommand(CmdClaimRewards())
	cmd.AddCommand(CmdRequestWithdrawAll())
	// this line is used by starport scaffolding # 1

	return cmd
}

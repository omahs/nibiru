package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/spf13/cobra"

	"github.com/NibiruChain/nibiru/x/sudo/pb"
)

// GetTxCmd returns a cli command for this module's transactions
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        pb.ModuleName,
		Short:                      fmt.Sprintf("x/%s transaction subcommands", pb.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Add subcommands
	txCmd.AddCommand(
		CmdEditSudoers(),
	)

	return txCmd
}

// GetQueryCmd returns a cli command for this module's queries
func GetQueryCmd() *cobra.Command {
	moduleQueryCmd := &cobra.Command{
		Use: pb.ModuleName,
		Short: fmt.Sprintf(
			"Query commands for the x/%s module", pb.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Add subcommands
	cmds := []*cobra.Command{
		CmdQuerySudoers(),
	}
	for _, cmd := range cmds {
		moduleQueryCmd.AddCommand(cmd)
	}

	return moduleQueryCmd
}

// CmdEditSudoers is a terminal command corresponding to the EditSudoers
// function of the sdk.Msg handler for x/sudo.
func CmdEditSudoers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [edit-json]",
		Args:  cobra.ExactArgs(1),
		Short: "Edit the x/sudo state (sudoers) by adding or removing contracts",
		Example: strings.TrimSpace(fmt.Sprintf(`
			Example: 
			$ %s tx sudo edit <path/to/edit.json> --from=<key_or_address> 
			`, version.AppName)),
		Long: strings.TrimSpace(
			`Adds or removes contracts from the x/sudo state, giving the 
			contracts permissioned access to certain bindings in x/wasm.

			The edit.json for 'EditSudoers' is of the form:
			{
			  "action": "add_contracts",
			  "contracts": "..."
			}

			- Valid action types: "add_contracts", "remove_contracts"	
			`),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := new(pb.MsgEditSudoers)

			// marshals contents into the proto.Message to which 'msg' points.
			contents, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			if err = clientCtx.Codec.UnmarshalJSON(contents, msg); err != nil {
				return err
			}

			// Parse the message sender
			from := clientCtx.GetFromAddress()
			msg.Sender = from.String()

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdQuerySudoers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state",
		Short: "displays the internal state (sudoers) of the x/sudo module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := pb.NewQueryClient(clientCtx)

			req := new(pb.QuerySudoersRequest)
			resp, err := queryClient.QuerySudoers(
				cmd.Context(), req,
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

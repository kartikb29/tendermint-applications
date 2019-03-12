package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdCreateRecord is the CLI command for sending a CreateRecord transaction
func GetCmdCreateRecord(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-record [_id] [data]",
		Short: "create a new record in the distributed storage engine with an identifier _id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// coins, err := sdk.ParseCoins(args[1])
			// if err != nil {
			// 	return err
			// }

			msg := datastore.NewMsgCreateRecord(args[0], args[1], cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

// GetCmdModifyRecordData is the CLI command for sending a ModifyRecord transaction
func GetCmdModifyRecordData(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "modify-record-data [_id] [data]",
		Short: "modify the data string of a record with an identifier _id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// coins, err := sdk.ParseCoins(args[1])
			// if err != nil {
			// 	return err
			// }

			msg := datastore.NewMsgModifyRecordData(args[0], args[1], cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

// GetCmdModifyRecordOwner is the CLI command for sending a ModifyOwner transaction
func GetCmdModifyRecordOwner(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "modify-record-owner [_id] [owner_accAddress]",
		Short: "modify the owner of a record with an identifier _id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// coins, err := sdk.ParseCoins(args[1])
			// if err != nil {
			// 	return err
			// }

			msg := datastore.NewMsgModifyRecordOwner(args[0], args[1], cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

// // GetCmdSetName is the CLI command for sending a SetName transaction
// func GetCmdSetName(cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "set-name [name] [value]",
// 		Short: "set the value associated with a name that you own",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

// 			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

// 			if err := cliCtx.EnsureAccountExists(); err != nil {
// 				return err
// 			}

// 			msg := nameservice.NewMsgSetName(args[0], args[1], cliCtx.GetFromAddress())
// 			err := msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			cliCtx.PrintResponse = true

// 			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
// 			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
// 		},
// 	}
// }

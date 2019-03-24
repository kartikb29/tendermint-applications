package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	datastore "github.com/kartikeya95/distributed-datastore/x/datastore"

	"github.com/spf13/cobra"
)

// GetCmdQueryRecord queries information about a record
func GetCmdQueryRecord(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "fetch [record_id]",
		Short: "fetch record_id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			_id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/record/%s", queryRoute, _id), nil)
			if err != nil {
				fmt.Printf("could not fetch data for ID - %s, ERR: %v \n", string(_id), err)
				return nil
			}

			var out datastore.QueryResData
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// // GetCmdWhois queries information about a domain
// func GetCmdWhois(queryRoute string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "whois [name]",
// 		Short: "Query whois info of name",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			name := args[0]

// 			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, name), nil)
// 			if err != nil {
// 				fmt.Printf("could not resolve whois - %s \n", string(name))
// 				return nil
// 			}

// 			var out nameservice.Whois
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return cliCtx.PrintOutput(out)
// 		},
// 	}
// }

// GetCmdRecords queries a list of all names
func GetCmdRecords(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "records",
		Short: "records",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/records", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not fetch records\n")
				return nil
			}

			var out datastore.QueryResRecords
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

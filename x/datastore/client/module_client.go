package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	datastorecmd "github.com/kartikeya95/distributed-datastore/x/datastore/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

//NewModuleClient - constructor function for the module client
func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group nameservice queries under a subcommand
	datastoreQueryCmd := &cobra.Command{
		Use:   "datastore",
		Short: "Querying commands for the datastore module",
	}

	datastoreQueryCmd.AddCommand(client.GetCommands(
		datastorecmd.GetCmdQueryRecord(mc.storeKey, mc.cdc),
		//datastoreQueryCmd.GetCmdWhois(mc.storeKey, mc.cdc),
		datastorecmd.GetCmdRecords(mc.storeKey, mc.cdc),
	)...)

	return datastoreQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	datatsoreTxCmd := &cobra.Command{
		Use:   "datatstore",
		Short: "datastore transactions subcommands",
	}

	datatsoreTxCmd.AddCommand(client.PostCommands(
		datastorecmd.GetCmdCreateRecord(mc.cdc),
		datastorecmd.GetCmdModifyRecordData(mc.cdc),
		datastorecmd.GetCmdModifyRecordOwner(mc.cdc),
	)...)

	return datatsoreTxCmd
}

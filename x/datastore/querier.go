package datastore

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the datastore Querier
const (
	QueryRecord = "record"
	//QueryRecordOwner = "record/owner"
	QueryRecords = "records"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryRecord:
			return queryRecord(ctx, path[1:], req, keeper)
		case QueryRecords:
			return queryRecords(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown datastore query endpoint")
		}
	}
}

// // nolint: unparam
// func queryRecord(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
// 	_id := path[0]

// 	data := keeper.GetData(ctx, _id)

// 	if data == "" {
// 		return []byte{}, sdk.ErrUnknownRequest("could not find record with that ID")
// 	}

// 	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResRecord{data})
// 	if err2 != nil {
// 		panic("could not marshal result to JSON")
// 	}

// 	return bz, nil
// }

// //QueryResResolve - Query Result Payload for a resolve query
// type QueryResResolve struct {
// 	Data string `json:"data"`
// }

// // implement fmt.Stringer
// func (r QueryResResolve) String() string {
// 	return r.Data
// }

// nolint: unparam
func queryRecord(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	_id := path[0]

	data := keeper.GetData(ctx, _id)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResData{data})
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

//QueryResData is struct for a fetch record query
type QueryResData struct {
	Data string `json:"Data"`
}

// implement fmt.Stringer
func (r QueryResData) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
	Data: %s`, r.Data))
}

func queryRecords(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var recordsList QueryResRecords

	iterator := keeper.GetRecordIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		record := string(iterator.Key())
		recordsList = append(recordsList, record)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, recordsList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// QueryResRecords - Query Result Payload for a names query
type QueryResRecords []string

// implement fmt.Stringer
func (n QueryResRecords) String() string {
	return strings.Join(n[:], "\n")
}

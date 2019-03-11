package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateRecord{}, "datastore/CreateRecord", nil)
	cdc.RegisterConcrete(MsgModifyRecordData{}, "datastore/ModifyRecordData", nil)
	cdc.RegisterConcrete(MsgModifyRecordOwner{}, "datastore/ModifyRecordOwner", nil)
}

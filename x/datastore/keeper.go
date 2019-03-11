package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	//"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	//coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the Record Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// GetRecord - gets the entire Record metadata struct for an _id
func (k Keeper) GetRecord(ctx sdk.Context, _id string) Record {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(_id)) {
		return NewRecord()
	}
	bz := store.Get([]byte(_id))
	var record Record
	k.cdc.MustUnmarshalBinaryBare(bz, &Record)
	return record
}

// SetRecord - sets the entire Record metadata struct for a _id
func (k Keeper) SetRecord(ctx sdk.Context, _id string, record Record) {
	if record.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(_id), k.cdc.MustMarshalBinaryBare(record))
}

// GetData - returns the data string that the _id resolves to
func (k Keeper) GetData(ctx sdk.Context, _id string) string {
	return k.GetRecord(ctx, _id).data
}

// SetData - sets the data string that a _id resolves to
func (k Keeper) SetData(ctx sdk.Context, _id string, data string) {
	record := k.GetRecord(ctx, _id)
	record.data = data
	k.SetRecord(ctx, _id, record)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, _id string) bool {
	return !k.GetRecord(ctx, _id).Owner.Empty()
}

// GetOwner - get the current owner of a record
func (k Keeper) GetOwner(ctx sdk.Context, _id string) sdk.AccAddress {
	return k.GetRecord(ctx, _id).Owner
}

// SetOwner - sets the current owner of a record
func (k Keeper) SetOwner(ctx sdk.Context, _id string, owner sdk.AccAddress) {
	record := k.GetRecord(ctx, _id)
	record.Owner = owner
	k.SetRecord(ctx, _id, record)
}

// // GetPrice - gets the current price of a name
// func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
// 	return k.GetWhois(ctx, name).Price
// }

// // SetPrice - sets the current price of a name
// func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
// 	whois := k.GetWhois(ctx, name)
// 	whois.Price = price
// 	k.SetWhois(ctx, name, whois)
// }

// GetRecordIterator - Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetRecordIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

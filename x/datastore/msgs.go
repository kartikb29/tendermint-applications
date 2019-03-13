package datastore

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreateRecord defines a CreateRecord message
type MsgCreateRecord struct {
	_id   string
	data  string
	Owner sdk.AccAddress
}

//MsgModifyRecordData defines a ModifyRecordOwner message
type MsgModifyRecordData struct {
	_id   string
	data  string
	Owner sdk.AccAddress
}

//MsgModifyRecordOwner defines a ModifyRecordOwner message
type MsgModifyRecordOwner struct {
	_id   string
	Owner sdk.AccAddress
}

// NewMsgCreateRecord is a constructor function for MsgSetRecord
func NewMsgCreateRecord(_id string, data string, owner sdk.AccAddress) MsgCreateRecord {
	return MsgCreateRecord{
		_id:   _id,
		data:  data,
		Owner: owner,
	}
}

//NewMsgModifyRecordData is s constructor function for MsgModifyRecordData
func NewMsgModifyRecordData(_id string, data string, owner sdk.AccAddress) MsgModifyRecordData {
	return MsgModifyRecordData{
		_id:   _id,
		data:  data,
		Owner: owner,
	}
}

//NewMsgModifyRecordOwner is s constructor function for MsgModifyRecordOwner
func NewMsgModifyRecordOwner(_id string, data string, owner sdk.AccAddress) MsgModifyRecordOwner {
	return MsgModifyRecordOwner{
		_id:   _id,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgCreateRecord) Route() string { return "datastore" }

// Type should return the action
func (msg MsgCreateRecord) Type() string { return "set_record" }

// Route should return the name of the module
func (msg MsgModifyRecordData) Route() string { return "datastore" }

// Type should return the action
func (msg MsgModifyRecordData) Type() string { return "set_data" }

// Route should return the name of the module
func (msg MsgModifyRecordOwner) Route() string { return "datastore" }

// Type should return the action
func (msg MsgModifyRecordOwner) Type() string { return "set_owner" }

//ValidateBasic runs stateless checks on the message
func (msg MsgCreateRecord) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg._id) == 0 || len(msg.data) == 0 {
		return sdk.ErrUnknownRequest("The _id or data cannot be empty")
	}
	return nil
}

//ValidateBasic runs stateless checks on the message
func (msg MsgModifyRecordData) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg._id) == 0 || len(msg.data) == 0 {
		return sdk.ErrUnknownRequest("The _id or data cannot be empty")
	}
	return nil
}

//ValidateBasic runs stateless checks on the message
func (msg MsgModifyRecordOwner) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg._id) == 0 {
		return sdk.ErrUnknownRequest("The _id or data cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRecord) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSignBytes encodes the message for signing
func (msg MsgModifyRecordData) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSignBytes encodes the message for signing
func (msg MsgModifyRecordOwner) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgCreateRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSigners defines whose signature is required
func (msg MsgModifyRecordData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSigners defines whose signature is required
func (msg MsgModifyRecordOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// // MsgBuyName defines the BuyName message
// type MsgBuyName struct {
// 	Name  string
// 	Bid   sdk.Coins
// 	Buyer sdk.AccAddress
// }

// // NewMsgBuyName is the constructor function for MsgBuyName
// func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
// 	return MsgBuyName{
// 		Name:  name,
// 		Bid:   bid,
// 		Buyer: buyer,
// 	}
// }

// // Route should return the name of the module
// func (msg MsgBuyName) Route() string { return "nameservice" }

// // Type should return the action
// func (msg MsgBuyName) Type() string { return "buy_name" }

// // ValidateBasic runs stateless checks on the message
// func (msg MsgBuyName) ValidateBasic() sdk.Error {
// 	if msg.Buyer.Empty() {
// 		return sdk.ErrInvalidAddress(msg.Buyer.String())
// 	}
// 	if len(msg.Name) == 0 {
// 		return sdk.ErrUnknownRequest("Name cannot be empty")
// 	}
// 	if !msg.Bid.IsAllPositive() {
// 		return sdk.ErrInsufficientCoins("Bids must be positive")
// 	}
// 	return nil
// }

// // GetSignBytes encodes the message for signing
// func (msg MsgBuyName) GetSignBytes() []byte {
// 	b, err := json.Marshal(msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return sdk.MustSortJSON(b)
// }

// // GetSigners defines whose signature is required
// func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
// 	return []sdk.AccAddress{msg.Buyer}
// }

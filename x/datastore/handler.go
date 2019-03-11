package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "datastore" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateRecord:
			return handleMsgCreateRecord(ctx, keeper, msg)
		case MsgModifyRecordData:
			return handleMsgModifyRecordData(ctx, keeper, msg)
		case MsgModifyRecordOwner:
			return handleMsgModifyRecordOwner(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

//Handles a message to create Record
func handleMsgCreateRecord(ctx sdk.Context, keeper Keeper, msg MsgCreateRecord) sdk.Result {
	if len(keeper.GetOwner(ctx, msg._id)) != 0 {
		errMsg := "A record with that ID already exists, please use a unique ID"
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
	keeper.SetOwner(ctx, msg._id, msg.Owner)
	keeper.SetData(ctx, msg._id, msg.data)
	return sdk.Result{}
}

// Handle a message to modify record data
func handleMsgModifyRecordData(ctx sdk.Context, keeper Keeper, msg MsgModifyRecordData) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg._id)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetData(ctx, msg._id, msg.data) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                    // return
}

// Handle a message to modify record owner
func handleMsgModifyRecordData(ctx sdk.Context, keeper Keeper, msg MsgModifyRecordData) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg._id)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetOwner(ctx, msg._id, msg.Owner) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}

// // Handle a message to buy name
// func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
// 	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) { // Checks if the the bid price is greater than the price paid by the current owner
// 		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
// 	}
// 	if keeper.HasOwner(ctx, msg.Name) {
// 		_, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
// 		if err != nil {
// 			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
// 		}
// 	} else {
// 		_, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
// 		if err != nil {
// 			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
// 		}
// 	}
// 	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
// 	keeper.SetPrice(ctx, msg.Name, msg.Bid)
// 	return sdk.Result{}
// }

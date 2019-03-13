package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	datastore "github.com/kartikeya95/datastore/x/datastore"

	"github.com/gorilla/mux"
)

// const (
// 	restName = "record"
// )

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, _id string) {
	r.HandleFunc(fmt.Sprintf("/datastore/record/%s", _id), createRecordHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/datastore/record/data/%s", _id), modifyRecordDataHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/datastore/record/owner/%s", _id), modifyRecordOwnerHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/datastore/record/%s", _id), getRecordHandler(cdc, cliCtx, _id)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/datastore/records"), recordsHandler(cdc, cliCtx)).Methods("GET")
}

type createRecordReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	_id     string       `json:"_id"`
	data    string       `json:"data"`
	Owner   string       `json:"Owner"`
}

func createRecordHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRecordReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// coins, err := sdk.ParseCoins(req.Amount)
		// if err != nil {
		// 	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// create the message
		msg := datastore.NewMsgCreateRecord(req._id, data, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)
	}
}

func modifyRecordDataHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRecordReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// coins, err := sdk.ParseCoins(req.Amount)
		// if err != nil {
		// 	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// create the message
		msg := datastore.NewMsgModifyRecordData(req._id, data, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)
	}
}

func modifyRecordOwnerHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRecordReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Buyer)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// coins, err := sdk.ParseCoins(req.Amount)
		// if err != nil {
		// 	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// create the message
		msg := datastore.NewMsgModifyRecordOwner(req._id, data, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)
	}
}

// type setNameReq struct {
// 	BaseReq rest.BaseReq `json:"base_req"`
// 	Name    string       `json:"name"`
// 	Value   string       `json:"value"`
// 	Owner   string       `json:"owner"`
// }

// func setNameHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req setNameReq
// 		if !rest.ReadRESTReq(w, r, cdc, &req) {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
// 			return
// 		}

// 		baseReq := req.BaseReq.Sanitize()
// 		if !baseReq.ValidateBasic(w) {
// 			return
// 		}

// 		addr, err := sdk.AccAddressFromBech32(req.Owner)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		// create the message
// 		msg := nameservice.NewMsgSetName(req.Name, req.Value, addr)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)
// 	}
// }

func getRecordHandler(cdc *codec.Codec, cliCtx context.CLIContext, _id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/record/data/%s", _id, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

// func whoIsHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		paramType := vars[restName]

// 		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", storeName, paramType), nil)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
// 			return
// 		}

// 		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
// 	}
// }

func recordsHandler(cdc *codec.Codec, cliCtx context.CLIContext, _id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/records", _id), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

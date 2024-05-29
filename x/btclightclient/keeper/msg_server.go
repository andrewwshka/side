package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sideprotocol/side/x/btclightclient/types"
)

type msgServer struct {
	Keeper
}

// SubmitBlockHeaders implements types.MsgServer.
func (m msgServer) SubmitBlockHeaders(goCtx context.Context, msg *types.MsgSubmitBlockHeaderRequest) (*types.MsgSubmitBlockHeadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// check if the sender is one of the authorized senders
	param := m.GetParams(ctx)
	if !param.IsAuthorizedSender(msg.Sender) {
		return nil, types.ErrSenderAddressNotAuthorized
	}

	// Set block headers
	err := m.SetBlockHeaders(ctx, msg.BlockHeaders)
	if err != nil {
		return nil, err
	}

	// Emit events
	// m.EmitEvent(
	// 	ctx,
	// 	msg.Sender,
	// 	sdk.Attribute{
	// 		Key:   types.AttributeKeyPoolCreator,
	// 		Value: msg.Sender,
	// 	},
	// )
	return &types.MsgSubmitBlockHeadersResponse{}, nil
}

// SubmitTransaction implements types.MsgServer.
// No Permission check required for this message
// Since everyone can submit a transaction to mint voucher tokens
// This message is usually sent by relayers
func (m msgServer) SubmitTransaction(goCtx context.Context, msg *types.MsgSubmitTransactionRequest) (*types.MsgSubmitTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		ctx.Logger().Error("Error validating basic", "error", err)
		return nil, err
	}

	if err := m.ProcessBitcoinDepositTransaction(ctx, msg); err != nil {
		ctx.Logger().Error("Error processing bitcoin deposit transaction", "error", err)
		return nil, err
	}

	// Emit Events
	m.EmitEvent(ctx, msg.Sender,
		sdk.NewAttribute("blockhash", msg.Blockhash),
		sdk.NewAttribute("txBytes", msg.TxBytes),
	)

	return &types.MsgSubmitTransactionResponse{}, nil

}

// UpdateSenders implements types.MsgServer.
func (m msgServer) UpdateSenders(goCtx context.Context, msg *types.MsgUpdateSendersRequest) (*types.MsgUpdateSendersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// check if the sender is one of the authorized senders
	param := m.GetParams(ctx)
	if !param.IsAuthorizedSender(msg.Sender) {
		return nil, types.ErrSenderAddressNotAuthorized
	}

	// Set block headers
	m.SetParams(ctx, types.NewParams(msg.Senders))

	// Emit events

	return &types.MsgUpdateSendersResponse{}, nil
}

func (m msgServer) WithdrawBitcoin(goCtx context.Context, msg *types.MsgWithdrawBitcoinRequest) (*types.MsgWithdrawBitcoinResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := sdk.AccAddress(msg.Sender)

	coin, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	balance := m.bankKeeper.GetBalance(ctx, sender, coin.Denom)

	if balance.Amount.LT(coin.Amount) {
		return nil, types.ErrInsufficientBalance
	}

	m.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(coin))

	//request := types.BitcoinSigningRequest(sender, coin)

	// Emit events
	m.EmitEvent(ctx, msg.Sender,
		sdk.NewAttribute("withdraw", msg.Amount),
	)

	return &types.MsgWithdrawBitcoinResponse{}, nil
}

// SubmitWithdrawSignatures submits the signatures of the withdraw transaction.
func (m msgServer) SubmitWithdrawSignatures(goCtx context.Context, msg *types.MsgSubmitWithdrawSignaturesRequest) (*types.MsgSubmitWithdrawSignaturesResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	request := m.GetSigningRequest(ctx, msg.Txid)
	// request.
	if ok := request.ValidateSignatures(msg.Signatures); !ok {
		return nil, types.ErrInvalidSignatures

	}
	return &types.MsgSubmitWithdrawSignaturesResponse{}, nil

}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
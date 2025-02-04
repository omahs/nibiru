package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v4/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"

	"github.com/NibiruChain/nibiru/app/ante"
)

type AnteHandlerOptions struct {
	sdkante.HandlerOptions
	IBCKeeper *ibckeeper.Keeper

	TxCounterStoreKey sdk.StoreKey
	WasmConfig        wasmtypes.WasmConfig
}

/*
	NewAnteHandler returns and AnteHandler that checks and increments sequence

numbers, checks signatures and account numbers, and deducts fees from the first signer.
*/
func NewAnteHandler(options AnteHandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.SigGasConsumer == nil {
		options.SigGasConsumer = sdkante.DefaultSigVerificationGasConsumer
	}
	if options.IBCKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "ibc keeper is required for AnteHandler")
	}

	anteDecorators := []sdk.AnteDecorator{
		sdkante.NewSetUpContextDecorator(),
		wasmkeeper.NewLimitSimulationGasDecorator(options.WasmConfig.SimulationGasLimit),
		wasmkeeper.NewCountTXDecorator(options.TxCounterStoreKey),
		sdkante.NewRejectExtensionOptionsDecorator(),
		sdkante.NewMempoolFeeDecorator(),
		sdkante.NewValidateBasicDecorator(),
		sdkante.NewTxTimeoutHeightDecorator(),
		sdkante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewPostPriceFixedPriceDecorator(),
		sdkante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		sdkante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper), // Replace fee ante from cosmos auth with a custom one.
		// SetPubKeyDecorator must be called before all signature verification decorators
		sdkante.NewSetPubKeyDecorator(options.AccountKeeper),
		sdkante.NewValidateSigCountDecorator(options.AccountKeeper),
		sdkante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		sdkante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		sdkante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

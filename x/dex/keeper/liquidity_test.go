package keeper_test

import (
	"testing"

	"github.com/MatrixDao/matrix/x/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetSetDenomLiquidity(t *testing.T) {
	app, ctx := testutil.NewMatrixApp()

	// Write to store
	app.DexKeeper.SetDenomLiquidity(ctx, "mtrx", sdk.NewInt(1000))

	// Read from store
	amount := app.DexKeeper.GetDenomLiquidity(ctx, "mtrx")

	require.EqualValues(t, sdk.NewInt(1000), amount)
}

func TestGetTotalLiquidity(t *testing.T) {
	app, ctx := testutil.NewMatrixApp()

	// Write to store
	app.DexKeeper.SetDenomLiquidity(ctx, "atom", sdk.NewInt(123))
	app.DexKeeper.SetDenomLiquidity(ctx, "mtrx", sdk.NewInt(456))
	app.DexKeeper.SetDenomLiquidity(ctx, "foo", sdk.NewInt(789))

	// Read from store
	coins := app.DexKeeper.GetTotalLiquidity(ctx)

	require.EqualValues(t, sdk.NewCoins(
		sdk.NewCoin("atom", sdk.NewInt(123)),
		sdk.NewCoin("mtrx", sdk.NewInt(456)),
		sdk.NewCoin("foo", sdk.NewInt(789)),
	), coins)
}

func TestSetTotalLiquidity(t *testing.T) {
	app, ctx := testutil.NewMatrixApp()

	// Write to store
	app.DexKeeper.SetTotalLiquidity(ctx, sdk.NewCoins(
		sdk.NewCoin("atom", sdk.NewInt(123)),
		sdk.NewCoin("mtrx", sdk.NewInt(456)),
		sdk.NewCoin("foo", sdk.NewInt(789)),
	))

	// Read from store
	require.EqualValues(t, app.DexKeeper.GetDenomLiquidity(ctx, "atom"), sdk.NewInt(123))
	require.EqualValues(t, app.DexKeeper.GetDenomLiquidity(ctx, "mtrx"), sdk.NewInt(456))
	require.EqualValues(t, app.DexKeeper.GetDenomLiquidity(ctx, "foo"), sdk.NewInt(789))
}

func TestRecordTotalLiquidityIncrease(t *testing.T) {
	app, ctx := testutil.NewMatrixApp()

	// Write to store
	app.DexKeeper.SetTotalLiquidity(ctx, sdk.NewCoins(
		sdk.NewCoin("atom", sdk.NewInt(100)),
		sdk.NewCoin("mtrx", sdk.NewInt(200)),
	))
	app.DexKeeper.RecordTotalLiquidityIncrease(ctx, sdk.NewCoins(
		sdk.NewCoin("atom", sdk.NewInt(50)),
		sdk.NewCoin("mtrx", sdk.NewInt(75)),
	))

	require.EqualValues(t, app.DexKeeper.GetDenomLiquidity(ctx, "atom"), sdk.NewInt(150))
	require.EqualValues(t, app.DexKeeper.GetDenomLiquidity(ctx, "mtrx"), sdk.NewInt(275))
}

func TestRecordTotalLiquidityDecrease(t *testing.T) {
	app, ctx := testutil.NewMatrixApp()

	// Write to store
	app.DexKeeper.SetTotalLiquidity(ctx, sdk.NewCoins(
		sdk.NewCoin("atom", sdk.NewInt(100)),
		sdk.NewCoin("mtrx", sdk.NewInt(200)),
	))
	app.DexKeeper.RecordTotalLiquidityDecrease(ctx, sdk.NewCoins(
		sdk.NewCoin("atom", sdk.NewInt(50)),
		sdk.NewCoin("mtrx", sdk.NewInt(75)),
	))

	require.EqualValues(t, app.DexKeeper.GetDenomLiquidity(ctx, "atom"), sdk.NewInt(50))
	require.EqualValues(t, app.DexKeeper.GetDenomLiquidity(ctx, "mtrx"), sdk.NewInt(125))
}
package binding_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/NibiruChain/collections"

	"github.com/NibiruChain/nibiru/app"
	"github.com/NibiruChain/nibiru/x/common/asset"
	"github.com/NibiruChain/nibiru/x/common/denoms"
	"github.com/NibiruChain/nibiru/x/common/testutil"
	"github.com/NibiruChain/nibiru/x/common/testutil/testapp"
	"github.com/NibiruChain/nibiru/x/wasm/binding"
	"github.com/NibiruChain/nibiru/x/wasm/binding/cw_struct"
	"github.com/NibiruChain/nibiru/x/wasm/binding/wasmbin"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

func TestSuitePerpExecutor_RunAll(t *testing.T) {
	suite.Run(t, new(TestSuitePerpExecutor))
}

type TestSuitePerpExecutor struct {
	suite.Suite

	nibiru           *app.NibiruApp
	ctx              sdk.Context
	contractDeployer sdk.AccAddress
	exec             *binding.ExecutorPerp

	contractPerp sdk.AccAddress
	ratesMap     map[asset.Pair]sdk.Dec
	happyFields  ExampleFields
}

func (s *TestSuitePerpExecutor) SetupSuite() {
	s.happyFields = GetHappyFields()
	sender := testutil.AccAddress()
	s.contractDeployer = sender

	genesisState := SetupPerpGenesis()
	nibiru := testapp.NewNibiruTestApp(genesisState)
	ctx := nibiru.NewContext(false, tmproto.Header{
		Height:  1,
		ChainID: "nibiru-wasmnet-1",
		Time:    time.Now().UTC(),
	})
	coins := sdk.NewCoins(
		sdk.NewCoin(denoms.NIBI, sdk.NewInt(1_000_000)),
		sdk.NewCoin(denoms.NUSD, sdk.NewInt(420_000*69)),
	)
	s.NoError(testapp.FundAccount(nibiru.BankKeeper, ctx, sender, coins))

	nibiru, ctx = SetupAllContracts(s.T(), sender, nibiru, ctx)
	s.nibiru = nibiru
	s.ctx = ctx

	wasmkeeper.NewMsgServerImpl(wasmkeeper.NewDefaultPermissionKeeper(nibiru.WasmKeeper))
	s.contractPerp = ContractMap[wasmbin.WasmKeyPerpBinding]
	s.exec = &binding.ExecutorPerp{
		Perp: nibiru.PerpKeeper,
	}
	s.OnSetupEnd()
}

func (s *TestSuitePerpExecutor) OnSetupEnd() {
	s.ratesMap = SetExchangeRates(s.Suite, s.nibiru, s.ctx)
}

// Happy path coverage of OpenPosition, AddMargin, RemoveMargin, and ClosePosition
func (s *TestSuitePerpExecutor) TestOpenAddRemoveClose() {
	pair := asset.MustNewPair(s.happyFields.Pair)
	margin := sdk.NewCoin(denoms.NUSD, sdk.NewInt(69))

	for _, err := range []error{
		s.DoOpenPositionTest(pair),
		s.DoAddMarginTest(pair, margin),
		s.DoRemoveMarginTest(pair, margin),
		s.DoClosePositionTest(pair),
		s.DoPegShiftTest(pair),
	} {
		s.NoError(err)
	}
}

func (s *TestSuitePerpExecutor) DoOpenPositionTest(pair asset.Pair) error {
	cwMsg := &cw_struct.OpenPosition{
		Sender:          s.contractDeployer.String(),
		Pair:            pair.String(),
		IsLong:          false,
		QuoteAmount:     sdk.NewInt(4_200_000),
		Leverage:        sdk.NewDec(5),
		BaseAmountLimit: sdk.NewInt(0),
	}

	_, err := s.exec.OpenPosition(cwMsg, s.ctx)
	if err != nil {
		return err
	}

	// Verify position exists with PerpKeeper
	_, err = s.exec.Perp.Positions.Get(
		s.ctx, collections.Join(pair, s.contractDeployer),
	)
	if err != nil {
		return err
	}

	// Verify position exists with CustomQuerier - multi-position
	bindingQuery := cw_struct.BindingQuery{
		Positions: &cw_struct.PositionsRequest{
			Trader: s.contractDeployer.String(),
		},
	}
	bindingRespMulti := new(cw_struct.PositionsRequest)
	_, err = DoCustomBindingQuery(
		s.ctx, s.nibiru, s.contractPerp, bindingQuery, bindingRespMulti,
	)
	if err != nil {
		return err
	}

	// Verify position exists with CustomQuerier - single position
	bindingQuery = cw_struct.BindingQuery{
		Position: &cw_struct.PositionRequest{
			Trader: s.contractDeployer.String(),
			Pair:   pair.String(),
		},
	}
	bindingResp := new(cw_struct.PositionRequest)
	_, err = DoCustomBindingQuery(
		s.ctx, s.nibiru, s.contractPerp, bindingQuery, bindingResp,
	)

	return err
}

func (s *TestSuitePerpExecutor) DoAddMarginTest(
	pair asset.Pair, margin sdk.Coin) error {
	cwMsg := &cw_struct.AddMargin{
		Sender: s.contractDeployer.String(),
		Pair:   pair.String(),
		Margin: margin,
	}

	_, err := s.exec.AddMargin(cwMsg, s.ctx)
	return err
}

func (s *TestSuitePerpExecutor) DoRemoveMarginTest(
	pair asset.Pair, margin sdk.Coin) error {
	cwMsg := &cw_struct.RemoveMargin{
		Sender: s.contractDeployer.String(),
		Pair:   pair.String(),
		Margin: margin,
	}

	_, err := s.exec.RemoveMargin(cwMsg, s.ctx)
	return err
}

func (s *TestSuitePerpExecutor) DoClosePositionTest(pair asset.Pair) error {
	cwMsg := &cw_struct.ClosePosition{
		Sender: s.contractDeployer.String(),
		Pair:   pair.String(),
	}

	_, err := s.exec.ClosePosition(cwMsg, s.ctx)
	return err
}

func (s *TestSuitePerpExecutor) DoPegShiftTest(pair asset.Pair) error {
	contractAddr := s.contractPerp
	cwMsg := &cw_struct.PegShift{
		Pair:    pair.String(),
		PegMult: sdk.NewDec(420),
	}

	err := s.exec.PegShift(cwMsg, contractAddr, s.ctx)
	return err
}

func (s *TestSuitePerpExecutor) DoDepthShiftTest(pair asset.Pair) error {
	contractAddr := s.contractPerp
	cwMsg := &cw_struct.DepthShift{
		Pair:      pair.String(),
		DepthMult: sdk.NewDec(420),
	}

	err := s.exec.DepthShift(cwMsg, contractAddr, s.ctx)
	return err
}

func (s *TestSuitePerpExecutor) TestSadPaths_Nil() {
	var err error

	_, err = s.exec.OpenPosition(nil, s.ctx)
	s.Error(err)

	_, err = s.exec.AddMargin(nil, s.ctx)
	s.Error(err)

	_, err = s.exec.RemoveMargin(nil, s.ctx)
	s.Error(err)

	_, err = s.exec.ClosePosition(nil, s.ctx)
	s.Error(err)

	err = s.exec.PegShift(
		nil, sdk.AccAddress([]byte("contract")), s.ctx)
	s.Error(err)

	err = s.exec.DepthShift(
		nil, sdk.AccAddress([]byte("contract")), s.ctx)
	s.Error(err)
}

func (s *TestSuitePerpExecutor) TestSadPaths_InvalidPair() {
	sadPair := asset.Pair("ftt:ust:doge")
	pair := sadPair
	margin := sdk.NewCoin(denoms.NUSD, sdk.NewInt(69))

	for _, err := range []error{
		s.DoOpenPositionTest(pair),
		s.DoAddMarginTest(pair, margin),
		s.DoRemoveMarginTest(pair, margin),
		s.DoClosePositionTest(pair),
		s.DoPegShiftTest(pair),
		s.DoDepthShiftTest(pair),
	} {
		s.Error(err)
	}
}

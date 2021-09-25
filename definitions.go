// Package okex_v5_go_sdk is generally a golang Api wrapper of Okex V5 API
//
// https://www.okex.com/docs-v5/en
package okex

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type (
	BaseUrl        string
	InstrumentType string
	MarginMode     string
	ContractType   string
	PositionType   string
	PositionSide   string
	TradeMode      string
	CountAction    string
	OrderSide      string
	GreekType      string
	BarSize        string
	TradeSide      string
	ChannelName    string
	Operation      string
	EventType      string
	OrderType      string
	QuantityType   string
	OrderFlowType  string
	OrderState     string
	ActionType     string
	APIKeyAccess   string

	Destination           int
	BillType              uint8
	BillSubType           uint8
	FeeCategory           uint8
	TransferType          uint8
	AccountType           uint8
	DepositState          uint8
	WithdrawalDestination uint8
	WithdrawalState       int8

	JsonTime time.Time

	ClientError error
)

const (
	RestUrl      = BaseUrl("https://www.okex.com")
	PublicWsUrl  = BaseUrl("wss://ws.okex.com:8443/ws/public")
	PrivateWsUrl = BaseUrl("wss://ws.okex.com:8443/ws/private")

	AwsRestUrl      = BaseUrl("https://aws.okex.com")
	AwsPublicWsUrl  = BaseUrl("wss://wsaws.okex.com:8443/ws/public")
	AwsPrivateWsUrl = BaseUrl("wss://wsaws.okex.com:8443/ws/private")

	DemoRestUrl      = BaseUrl("https://www.okex.com")
	DemoPublicWsUrl  = BaseUrl("wss://wspap.okex.com:8443/ws/public?brokerId=9999")
	DemoPrivateWsUrl = BaseUrl("wss://wspap.okex.com:8443/ws/private?brokerId=9999")

	NormalServer = Destination(iota + 1)
	AwsServer    = NormalServer + 1
	DemoServer   = AwsServer + 1

	SpotInstrument    = InstrumentType("SPOT")
	MarginInstrument  = InstrumentType("MARGIN")
	SwapInstrument    = InstrumentType("SWAP")
	FuturesInstrument = InstrumentType("FUTURES")
	OptionsInstrument = InstrumentType("OPTION")

	MarginCrossMode    = MarginMode("cross")
	MarginIsolatedMode = MarginMode("isolated")

	ContractLinearType  = ContractType("linear")
	ContractInverseType = ContractType("inverse")

	BillTransferType              = BillType(1)
	BillTradeType                 = BillType(2)
	BillDeliveryType              = BillType(3)
	BillAutoTokenConversionType   = BillType(4)
	BillLiquidationType           = BillType(5)
	BillMarginTransferType        = BillType(6)
	BillInterestDeductionType     = BillType(7)
	BillFundingFeeType            = BillType(8)
	BillADLType                   = BillType(9)
	BillClawbackType              = BillType(10)
	BillSystemTokenConversionType = BillType(11)
	BillStrategyTransferType      = BillType(12)

	BillBuySubType                              = BillSubType(1)
	BillSellSubType                             = BillSubType(2)
	BillOpenLongSubType                         = BillSubType(3)
	BillOpenShortSubType                        = BillSubType(4)
	BillCloseLongSubType                        = BillSubType(5)
	BillCloseShortSubType                       = BillSubType(6)
	BillInterestDeductionSubType                = BillSubType(9)
	BillTransferInSubType                       = BillSubType(11)
	BillTransferOutSubType                      = BillSubType(12)
	BillManualMarginIncreaseSubType             = BillSubType(160)
	BillManualMarginDecreaseSubType             = BillSubType(161)
	BillAutoMarginIncreaseSubType               = BillSubType(162)
	BillAutoBuySubType                          = BillSubType(110)
	BillAutoSellSubType                         = BillSubType(111)
	BillSystemTokenConversionTransferInSubType  = BillSubType(118)
	BillSystemTokenConversionTransferOutSubType = BillSubType(119)
	BillPartialLiquidationCloseLongSubType      = BillSubType(100)
	BillPartialLiquidationCloseShortSubType     = BillSubType(101)
	BillPartialLiquidationBuySubType            = BillSubType(102)
	BillPartialLiquidationSellSubType           = BillSubType(103)
	BillLiquidationLongSubType                  = BillSubType(104)
	BillLiquidationShortSubType                 = BillSubType(105)
	BillLiquidationBuySubType                   = BillSubType(106)
	BillLiquidationSellSubType                  = BillSubType(107)
	BillLiquidationTransferInSubType            = BillSubType(110)
	BillLiquidationTransferOutSubType           = BillSubType(111)
	BillADLCloseLongSubType                     = BillSubType(125)
	BillADLCloseShortSubType                    = BillSubType(126)
	BillADLBuySubType                           = BillSubType(127)
	BillADLSellSubType                          = BillSubType(128)
	BillExercisedSubType                        = BillSubType(170)
	BillCounterpartyExercisedSubType            = BillSubType(171)
	BillExpiredOTMSubType                       = BillSubType(172)
	BillDeliveryLongSubType                     = BillSubType(112)
	BillDeliveryShortSubType                    = BillSubType(113)
	BillDeliveryExerciseClawbackSubType         = BillSubType(117)
	BillFundingFeeExpenseSubType                = BillSubType(173)
	BillFundingFeeIncomeSubType                 = BillSubType(174)
	BillSystemTransferInSubType                 = BillSubType(200)
	BillManuallyTransferInSubType               = BillSubType(201)
	BillSystemTransferOutSubType                = BillSubType(202)
	BillManuallyTransferOutSubType              = BillSubType(203)

	PositionLongShortMode = PositionType("long_short_mode")
	PositionNetMode       = PositionType("net_mode")

	PositionLongSide  = PositionSide("long")
	PositionShortSide = PositionSide("short")
	PositionNetSide   = PositionSide("net")

	TradeCrossMode    = TradeMode("cross")
	TradeIsolatedMode = TradeMode("isolated")
	TradeCashMode     = TradeMode("cash")

	CountIncrease = CountAction("add")
	CountDecrease = CountAction("reduce")

	OrderBuy  = OrderSide("buy")
	OrderSell = OrderSide("sell")

	GreekInCoin    = GreekType("PA")
	GreekInDollars = GreekType("PB")

	Bar1m  = BarSize("1m")
	Bar3m  = BarSize("3m")
	Bar5m  = BarSize("5m")
	Bar15m = BarSize("15m")
	Bar30m = BarSize("130")
	Bar1H  = BarSize("1H")
	Bar2H  = BarSize("2H")
	Bar4H  = BarSize("4H")
	Bar6H  = BarSize("6H")
	Bar8H  = BarSize("8H")
	Bar12H = BarSize("12H")
	Bar1D  = BarSize("1D")
	Bar1W  = BarSize("1W")
	Bar1M  = BarSize("1M")
	Bar3M  = BarSize("3M")
	Bar6M  = BarSize("6M")
	Bar1Y  = BarSize("1Y")

	TradeBuySide  = TradeSide("buy")
	TradeSellSide = TradeSide("sell")

	LoginOperation       = Operation("login")
	SubscribeOperation   = Operation("subscribe")
	UnsubscribeOperation = Operation("unsubscribe")

	OrderMarket          = OrderType("market")
	OrderLimit           = OrderType("limit")
	OrderPostOnly        = OrderType("post_only")
	OrderFOK             = OrderType("fok")
	OrderIOC             = OrderType("ioc")
	OrderOptimalLimitIoc = OrderType("optimal_limit_ioc")

	QuantityBaseCcy  = QuantityType("base_ccy")
	QuantityQuoteCcy = QuantityType("quote_ccy")

	OrderTakerFlow = OrderFlowType("T")
	OrderMakerFlow = OrderFlowType("M")

	ClassA = FeeCategory(1)
	ClassB = FeeCategory(2)
	ClassC = FeeCategory(3)
	ClassD = FeeCategory(4)

	OrderCancel          = OrderState("canceled")
	OrderLive            = OrderState("live")
	OrderPartiallyFilled = OrderState("partially_filled")
	OrderFilled          = OrderState("filled")

	TransferWithinAccount     = TransferType(0)
	MasterAccountToSubAccount = TransferType(1)
	MasterSubAccountToAccount = TransferType(2)

	SpotAccount    = AccountType(1)
	FuturesAccount = AccountType(3)
	MarginAccount  = AccountType(5)
	FundingAccount = AccountType(6)
	SwapAccount    = AccountType(9)
	OptionsAccount = AccountType(12)
	UnifiedAccount = AccountType(18)

	WaitingForConfirmation     = DepositState(0)
	DepositCredited            = DepositState(1)
	DepositSuccessful          = DepositState(2)
	DepositTemporarySuspension = DepositState(8)

	WithdrawalOkexDestination           = WithdrawalDestination(3)
	WithdrawalDigitalAddressDestination = WithdrawalDestination(4)

	WithdrawalPendingCancel              = WithdrawalState(-3)
	WithdrawalCanceled                   = WithdrawalState(-2)
	WithdrawalFailed                     = WithdrawalState(-1)
	WithdrawalPending                    = WithdrawalState(0)
	WithdrawalSending                    = WithdrawalState(1)
	WithdrawalSent                       = WithdrawalState(2)
	WithdrawalAwaitingEmailVerification  = WithdrawalState(3)
	WithdrawalAwaitingManualVerification = WithdrawalState(4)
	WithdrawalIdentityManualVerification = WithdrawalState(5)

	ActionPurchase = ActionType("purchase")
	ActionRedempt  = ActionType("redempt")

	APIKeyReadOnly = APIKeyAccess("read_only")
	APIKeyTrade    = APIKeyAccess("trade")
)

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *JsonTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.UnixMilli(q)
	return
}

func (t JsonTime) String() string { return time.Time(t).String() }

func (s BarSize) Duration() time.Duration {
	switch s {
	case Bar3m:
		return time.Minute * 3
	case Bar5m:
		return time.Minute * 5
	case Bar15m:
		return time.Minute * 15
	case Bar30m:
		return time.Minute * 30
	case Bar1H:
		return time.Hour * 1
	case Bar2H:
		return time.Hour * 2
	case Bar4H:
		return time.Hour * 4
	case Bar6H:
		return time.Hour * 6
	case Bar8H:
		return time.Hour * 8
	case Bar12H:
		return time.Hour * 12
	case Bar1D:
		return time.Hour * 24
	case Bar1W:
		return time.Hour * 24 * 7
	case Bar1M:
		return time.Hour * 24 * 30
	case Bar3M:
		return time.Hour * 24 * 30 * 3
	case Bar6M:
		return time.Hour * 20 * 30 * 6
	case Bar1Y:
		return time.Hour * 24 * 365
	}

	return time.Minute
}

func S2M(i interface{}) map[string]string {
	m := make(map[string]string)
	j, _ := json.Marshal(i)
	_ = json.Unmarshal(j, &m)

	return m
}

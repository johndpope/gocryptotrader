package bitfinex

// Ticker holds basic ticker information from the exchange
type Ticker struct {
	Mid       float64 `json:"mid,string"`
	Bid       float64 `json:"bid,string"`
	Ask       float64 `json:"ask,string"`
	Last      float64 `json:"last_price,string"`
	Low       float64 `json:"low,string"`
	High      float64 `json:"high,string"`
	Volume    float64 `json:"volume,string"`
	Timestamp string  `json:"timestamp"`
	Message   string  `json:"message"`
}

// Tickerv2 holds the version 2 ticker information
type Tickerv2 struct {
	FlashReturnRate float64
	Bid             float64
	BidPeriod       int64
	BidSize         float64
	Ask             float64
	AskPeriod       int64
	AskSize         float64
	DailyChange     float64
	DailyChangePerc float64
	Last            float64
	Volume          float64
	High            float64
	Low             float64
}

// Tickersv2 holds the version 2 tickers information
type Tickersv2 struct {
	Symbol string
	Tickerv2
}

// Stat holds individual statistics from exchange
type Stat struct {
	Period int64   `json:"period"`
	Volume float64 `json:"volume,string"`
}

// FundingBook holds current the full margin funding book
type FundingBook struct {
	Bids    []Book `json:"bids"`
	Asks    []Book `json:"asks"`
	Message string `json:"message"`
}

// Orderbook holds orderbook information from bid and ask sides
type Orderbook struct {
	Bids []Book
	Asks []Book
}

// BookV2 holds the orderbook item
type BookV2 struct {
	Price  float64
	Rate   float64
	Period float64
	Count  int64
	Amount float64
}

// OrderbookV2 holds orderbook information from bid and ask sides
type OrderbookV2 struct {
	Bids []BookV2
	Asks []BookV2
}

// TradeStructure holds executed trade information
type TradeStructure struct {
	Timestamp int64   `json:"timestamp"`
	Tid       int64   `json:"tid"`
	Price     float64 `json:"price,string"`
	Amount    float64 `json:"amount,string"`
	Exchange  string  `json:"exchange"`
	Type      string  `json:"sell"`
}

// TradeStructureV2 holds resp information
type TradeStructureV2 struct {
	Timestamp int64
	TID       int64
	Price     float64
	Amount    float64
	Exchange  string
	Type      string
}

// Lendbook holds most recent funding data for a relevant currency
type Lendbook struct {
	Bids []Book `json:"bids"`
	Asks []Book `json:"asks"`
}

// Book is a generalised sub-type to hold book information
type Book struct {
	Price           float64 `json:"price,string"`
	Rate            float64 `json:"rate,string"`
	Amount          float64 `json:"amount,string"`
	Period          int     `json:"period"`
	Timestamp       string  `json:"timestamp"`
	FlashReturnRate string  `json:"frr"`
}

// Lends holds the lent information by currency
type Lends struct {
	Rate       float64 `json:"rate,string"`
	AmountLent float64 `json:"amount_lent,string"`
	AmountUsed float64 `json:"amount_used,string"`
	Timestamp  int64   `json:"timestamp"`
}

// SymbolDetails holds currency pair information
type SymbolDetails struct {
	Pair             string  `json:"pair"`
	PricePrecision   int     `json:"price_precision"`
	InitialMargin    float64 `json:"initial_margin,string"`
	MinimumMargin    float64 `json:"minimum_margin,string"`
	MaximumOrderSize float64 `json:"maximum_order_size,string"`
	MinimumOrderSize float64 `json:"minimum_order_size,string"`
	Expiration       string  `json:"expiration"`
}

// AccountInfoFull adds the error message to Account info
type AccountInfoFull struct {
	Info    []AccountInfo
	Message string `json:"message"`
}

// AccountInfo general account information with fees
type AccountInfo struct {
	MakerFees string            `json:"maker_fees"`
	TakerFees string            `json:"taker_fees"`
	Fees      []AccountInfoFees `json:"fees"`
	Message   string            `json:"message"`
}

// AccountInfoFees general account information with fees
type AccountInfoFees struct {
	Pairs     string `json:"pairs"`
	MakerFees string `json:"maker_fees"`
	TakerFees string `json:"taker_fees"`
}

// AccountFees stores withdrawal account fee data from Bitfinex
type AccountFees struct {
	Withdraw struct {
		BTC  float64 `json:"BTC,string"`
		ETH  float64 `json:"ETH,string"`
		EOS  float64 `json:"EOS,string"`
		LTC  float64 `json:"LTC,string"`
		XRP  float64 `json:"XRP,string"`
		BCH  float64 `json:"BCH,string"`
		IOT  float64 `json:"IOT"`
		ETC  float64 `json:"ETC,string"`
		OMG  float64 `json:"OMG,string"`
		XMR  float64 `json:"XMR,string"`
		DSH  float64 `json:"DSH,string"`
		TRX  float64 `json:"TRX,string"`
		ZEC  float64 `json:"ZEC,string"`
		SAN  float64 `json:"SAN,string"`
		ZRX  float64 `json:"ZXR,string"`
		EDO  float64 `json:"EDO,string"`
		ETP  float64 `json:"ETP,string"`
		QTUM float64 `json:"QTUM,string"`
		BAT  float64 `json:"BAT,string"`
		SNT  float64 `json:"SNT,string"`
		AID  float64 `json:"AID,string"`
		RESP float64 `json:"QASH,string"`
		YYW  float64 `json:"YYW,string"`
		DATA float64 `json:"DATA,string"`
		GNT  float64 `json:"GNT,string"`
		TNB  float64 `json:"TNB,string"`
		FYN  float64 `json:"FYN,string"`
		SPK  float64 `json:"SPK,string"`
		RLC  float64 `json:"RLC,string"`
		ELF  float64 `json:"ELF,string"`
		AVT  float64 `json:"AVT,string"`
		MNA  float64 `json:"MNA,string"`
		RCN  float64 `json:"RCN,string"`
		SNG  float64 `json:"SNG,string"`
		NEC  float64 `json:"NEC,string"`
		IOS  float64 `json:"IOS,string"`
		AIO  float64 `json:"AIO,string"`
		REQ  float64 `json:"REQ,string"`
		RDN  float64 `json:"RDN,string"`
		LRC  float64 `json:"LRC,string"`
		WAX  float64 `json:"WAX,string"`
		DAI  float64 `json:"DAI,string"`
		CFI  float64 `json:"CFI,string"`
		AGI  float64 `json:"AGI,string"`
		BFT  float64 `json:"BFT,string"`
		MTN  float64 `json:"MTN,string"`
		ODE  float64 `json:"ODE,string"`
		ANT  float64 `json:"ANT,string"`
		DTH  float64 `json:"DTH,string"`
		MIT  float64 `json:"MIT,string"`
		STJ  float64 `json:"STJ,string"`
		XVG  float64 `json:"XVG,string"`
		MKR  float64 `json:"MKR,string"`
		VEN  float64 `json:"VEN,string"`
		KNC  float64 `json:"KNC,string"`
		POA  float64 `json:"POA,string"`
		EVT  float64 `json:"EVT,string"`
		LYM  float64 `json:"LYM,string"`
		UTK  float64 `json:"UTK,string"`
		VEE  float64 `json:"VEE,string"`
		DAD  float64 `json:"DAD,string"`
		ORS  float64 `json:"ORS,string"`
		AUC  float64 `json:"AUC,string"`
		POY  float64 `json:"POY,string"`
		FSN  float64 `json:"FSN,string"`
		CBT  float64 `json:"CBT,string"`
		ZCN  float64 `json:"ZCN,string"`
		SEN  float64 `json:"SEN,string"`
		NCA  float64 `json:"NCA,string"`
		CND  float64 `json:"CND,string"`
		CTX  float64 `json:"CTX,string"`
		PAI  float64 `json:"PAI,string"`
		SEE  float64 `json:"SEE,string"`
		ESS  float64 `json:"ESS,string"`
		MTO  float64 `json:"MTO,string"`
		ATM  float64 `json:"ATM,string"`
		HOT  float64 `json:"HOT,string"`
		DTA  float64 `json:"DTA,string"`
		IQX  float64 `json:"IQX,string"`
		WPR  float64 `json:"WPR,string"`
		ZIL  float64 `json:"ZIL,string"`
		BNT  float64 `json:"BNT,string"`
		ABS  float64 `json:"ABS,string"`
		XRA  float64 `json:"XRA,string"`
		MAN  float64 `json:"MAN,string"`
		BBN  float64 `json:"BBN,string"`
		NIO  float64 `json:"NIO,string"`
		DGX  float64 `json:"DGX,string"`
		UTN  float64 `json:"UTN,string"`
		TKN  float64 `json:"TKN,string"`
		GOT  float64 `json:"GOT,string"`
		XTZ  float64 `json:"XTZ,string"`
		CNN  float64 `json:"CNN,string"`
		BOX  float64 `json:"BOX,string"`
		USD  float64 `json:"USD,string"`
		EUR  float64 `json:"EUR,string"`
	} `json:"withdraw"`
}

// AccountSummary holds account summary data
type AccountSummary struct {
	TradeVolumePer30D []Currency `json:"trade_vol_30d"`
	FundingProfit30D  []Currency `json:"funding_profit_30d"`
	MakerFee          float64    `json:"maker_fee"`
	TakerFee          float64    `json:"taker_fee"`
}

// Currency is a sub-type for AccountSummary data
type Currency struct {
	Currency string  `json:"curr"`
	Volume   float64 `json:"vol,string"`
	Amount   float64 `json:"amount,string"`
}

// DepositResponse holds deposit address information
type DepositResponse struct {
	Result   string `json:"string"`
	Method   string `json:"method"`
	Currency string `json:"currency"`
	Address  string `json:"address"`
}

// KeyPermissions holds the key permissions for the API key set
type KeyPermissions struct {
	Account   Permission `json:"account"`
	History   Permission `json:"history"`
	Orders    Permission `json:"orders"`
	Positions Permission `json:"positions"`
	Funding   Permission `json:"funding"`
	Wallets   Permission `json:"wallets"`
	Withdraw  Permission `json:"withdraw"`
}

// Permission sub-type for KeyPermissions
type Permission struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

// MarginInfo holds metadata for margin information from bitfinex
type MarginInfo struct {
	Info    MarginData
	Message string `json:"message"`
}

// MarginData holds wallet information for margin trading
type MarginData struct {
	MarginBalance     float64        `json:"margin_balance,string"`
	TradableBalance   float64        `json:"tradable_balance,string"`
	UnrealizedPL      int64          `json:"unrealized_pl"`
	UnrealizedSwap    int64          `json:"unrealized_swap"`
	NetValue          float64        `json:"net_value,string"`
	RequiredMargin    int64          `json:"required_margin"`
	Leverage          float64        `json:"leverage,string"`
	MarginRequirement float64        `json:"margin_requirement,string"`
	MarginLimits      []MarginLimits `json:"margin_limits"`
}

// MarginLimits holds limit data per pair
type MarginLimits struct {
	OnPair            string  `json:"on_pair"`
	InitialMargin     float64 `json:"initial_margin,string"`
	MarginRequirement float64 `json:"margin_requirement,string"`
	TradableBalance   float64 `json:"tradable_balance,string"`
}

// Balance holds current balance data
type Balance struct {
	Type      string  `json:"type"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount,string"`
	Available float64 `json:"available,string"`
}

// WalletTransfer holds status of wallet to wallet content transfer on exchange
type WalletTransfer struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Withdrawal holds withdrawal status information
type Withdrawal struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	WithdrawalID int64  `json:"withdrawal_id,string"`
}

// Order holds order information when an order is in the market
type Order struct {
	ID                    int64   `json:"id"`
	Symbol                string  `json:"symbol"`
	Exchange              string  `json:"exchange"`
	Price                 float64 `json:"price,string"`
	AverageExecutionPrice float64 `json:"avg_execution_price,string"`
	Side                  string  `json:"side"`
	Type                  string  `json:"type"`
	Timestamp             string  `json:"timestamp"`
	IsLive                bool    `json:"is_live"`
	IsCancelled           bool    `json:"is_cancelled"`
	IsHidden              bool    `json:"is_hidden"`
	WasForced             bool    `json:"was_forced"`
	OriginalAmount        float64 `json:"original_amount,string"`
	RemainingAmount       float64 `json:"remaining_amount,string"`
	ExecutedAmount        float64 `json:"executed_amount,string"`
	OrderID               int64   `json:"order_id"`
}

// OrderMultiResponse holds order information on the executed orders
type OrderMultiResponse struct {
	Orders []Order `json:"order_ids"`
	Status string  `json:"status"`
}

// PlaceOrder is used for order placement
type PlaceOrder struct {
	Symbol   string  `json:"symbol"`
	Amount   float64 `json:"amount,string"`
	Price    float64 `json:"price,string"`
	Exchange string  `json:"exchange"`
	Side     string  `json:"side"`
	Type     string  `json:"type"`
}

// GenericResponse holds the result for a generic response
type GenericResponse struct {
	Result string `json:"result"`
}

// Position holds position information
type Position struct {
	ID        int64   `json:"id"`
	Symbol    string  `json:"string"`
	Status    string  `json:"active"`
	Base      float64 `json:"base,string"`
	Amount    float64 `json:"amount,string"`
	Timestamp string  `json:"timestamp"`
	Swap      float64 `json:"swap,string"`
	PL        float64 `json:"pl,string"`
}

// BalanceHistory holds balance history information
type BalanceHistory struct {
	Currency    string  `json:"currency"`
	Amount      float64 `json:"amount,string"`
	Balance     float64 `json:"balance,string"`
	Description string  `json:"description"`
	Timestamp   string  `json:"timestamp"`
}

// MovementHistory holds deposit and withdrawal history data
type MovementHistory struct {
	ID               int64   `json:"id"`
	TxID             int64   `json:"txid"`
	Currency         string  `json:"currency"`
	Method           string  `json:"method"`
	Type             string  `json:"withdrawal"`
	Amount           float64 `json:"amount,string"`
	Description      string  `json:"description"`
	Address          string  `json:"address"`
	Status           string  `json:"status"`
	Timestamp        string  `json:"timestamp"`
	TimestampCreated string  `json:"timestamp_created"`
	Fee              float64 `json:"fee"`
}

// TradeHistory holds trade history data
type TradeHistory struct {
	Price       float64 `json:"price,string"`
	Amount      float64 `json:"amount,string"`
	Timestamp   string  `json:"timestamp"`
	Exchange    string  `json:"exchange"`
	Type        string  `json:"type"`
	FeeCurrency string  `json:"fee_currency"`
	FeeAmount   float64 `json:"fee_amount,string"`
	TID         int64   `json:"tid"`
	OrderID     int64   `json:"order_id"`
}

// Offer holds offer information
type Offer struct {
	ID              int64   `json:"id"`
	Currency        string  `json:"currency"`
	Rate            float64 `json:"rate,string"`
	Period          int64   `json:"period"`
	Direction       string  `json:"direction"`
	Timestamp       string  `json:"timestamp"`
	Type            string  `json:"type"`
	IsLive          bool    `json:"is_live"`
	IsCancelled     bool    `json:"is_cancelled"`
	OriginalAmount  float64 `json:"original_amount,string"`
	RemainingAmount float64 `json:"remaining_amount,string"`
	ExecutedAmount  float64 `json:"executed_amount,string"`
}

// MarginFunds holds active funding information used in a margin position
type MarginFunds struct {
	ID         int64   `json:"id"`
	PositionID int64   `json:"position_id"`
	Currency   string  `json:"currency"`
	Rate       float64 `json:"rate,string"`
	Period     int     `json:"period"`
	Amount     float64 `json:"amount,string"`
	Timestamp  string  `json:"timestamp"`
	AutoClose  bool    `json:"auto_close"`
}

// MarginTotalTakenFunds holds position funding including sum of active backing
// as total swaps
type MarginTotalTakenFunds struct {
	PositionPair string  `json:"position_pair"`
	TotalSwaps   float64 `json:"total_swaps,string"`
}

// Fee holds fee data for a specified currency
type Fee struct {
	Currency  string
	TakerFees float64
	MakerFees float64
}

// WebsocketChanInfo holds websocket channel information
type WebsocketChanInfo struct {
	Channel string
	Pair    string
}

// WebsocketBook holds booking information
type WebsocketBook struct {
	Price  float64
	Count  int
	Amount float64
}

// WebsocketTrade holds trade information
type WebsocketTrade struct {
	ID        int64
	Timestamp int64
	Price     float64
	Amount    float64
}

// WebsocketTicker holds ticker information
type WebsocketTicker struct {
	Bid             float64
	BidSize         float64
	Ask             float64
	AskSize         float64
	DailyChange     float64
	DialyChangePerc float64
	LastPrice       float64
	Volume          float64
}

// WebsocketPosition holds position information
type WebsocketPosition struct {
	Pair              string
	Status            string
	Amount            float64
	Price             float64
	MarginFunding     float64
	MarginFundingType int
}

// WebsocketWallet holds wallet information
type WebsocketWallet struct {
	Name              string
	Currency          string
	Balance           float64
	UnsettledInterest float64
}

// WebsocketOrder holds order data
type WebsocketOrder struct {
	OrderID    int64
	Pair       string
	Amount     float64
	OrigAmount float64
	OrderType  string
	Status     string
	Price      float64
	PriceAvg   float64
	Timestamp  string
	Notify     int
}

// WebsocketTradeExecuted holds executed trade data
type WebsocketTradeExecuted struct {
	TradeID        int64
	Pair           string
	Timestamp      int64
	OrderID        int64
	AmountExecuted float64
	PriceExecuted  float64
}

// ErrorCapture is a simple type for returned errors from Bitfinex
type ErrorCapture struct {
	Message string `json:"message"`
}

// TimeInterval represents interval enum.
type TimeInterval string

// TimeInvterval vars
var (
	TimeIntervalMinute         = TimeInterval("1m")
	TimeIntervalFiveMinutes    = TimeInterval("5m")
	TimeIntervalFifteenMinutes = TimeInterval("15m")
	TimeIntervalThirtyMinutes  = TimeInterval("30m")
	TimeIntervalHour           = TimeInterval("1h")
	TimeIntervalThreeHours     = TimeInterval("3h")
	TimeIntervalSixHours       = TimeInterval("6h")
	TimeIntervalTwelveHours    = TimeInterval("12h")
	TimeIntervalDay            = TimeInterval("1d")
	TimeIntervalSevenDays      = TimeInterval("7d")
	TimeIntervalFourteenDays   = TimeInterval("14d")
	TimeIntervalMonth          = TimeInterval("1M")
)

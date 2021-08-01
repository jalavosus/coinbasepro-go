package endpoints

const (
	RestAPI       string = "https://api.pro.coinbase.com"
	WebsocketFeed string = "wss://ws-feed.pro.coinbase.com"
)

const (
	Accounts       string = "/accounts"
	Account               = Accounts + "/%[1]s" // formatted value must be an account Id.
	AccountHistory        = Account + "/ledger"
	AccountHolds          = Account + "/holds"
)

const (
	Orders string = "/orders"
	Order         = Orders + "/%[1]s" // formatted value must be an order Id.
)

const Fills string = "/fills"

const Transfers string = "/transfers"

const (
	transferTypeDeposit          string = "deposit"
	transferTypeInternalDeposit  string = "internal_deposit"
	transferTypeWithdrawal       string = "withdraw"
	transferTypeInternalWithdraw string = "internal_withdraw"
)

const (
	Products      string = "/products"
	Product              = Products + "/%[1]s" // formatted value must be a valid Product Id.
	ProductTicker        = Product + "/ticker"
	ProductTrades        = Product + "/trades"
)

const Currencies string = "/currencies"

const (
	UserExchangeLimits string = "/users/self/exchange-limits"
)

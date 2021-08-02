package endpoints

import (
	"github.com/jalavosus/coinbasepro-go/util/secrets"
)

const (
	ProductionRestAPIURI       string = "https://api.pro.coinbase.com"
	ProductionWebsocketFeedURI string = "wss://ws-feed.pro.coinbase.com"
)

const (
	SandboxRestAPIURI       string = "https://api-public.sandbox.pro.coinbase.com"
	SandboxWebsocketFeedURI string = "wss://ws-feed-public.sandbox.pro.coinbase.com"
)

func RestAPIURI() string {
	if secrets.UseSandbox() {
		return SandboxRestAPIURI
	}

	return ProductionRestAPIURI
}

func WebsocketFeedURI() string {
	if secrets.UseSandbox() {
		return SandboxWebsocketFeedURI
	}

	return ProductionWebsocketFeedURI
}

const (
	Accounts       string = "/accounts"
	Account               = Accounts + "/%[1]s" // formatted value must be an account ID.
	AccountHistory        = Account + "/ledger"
	AccountHolds          = Account + "/holds"
)

const (
	Orders string = "/orders"
	Order         = Orders + "/%[1]s" // formatted value must be an order ID.
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
	Product              = Products + "/%[1]s" // formatted value must be a valid Product ID.
	ProductTicker        = Product + "/ticker"
	ProductTrades        = Product + "/trades"
)

const Currencies string = "/currencies"

const (
	UserExchangeLimits string = "/users/self/exchange-limits"
)

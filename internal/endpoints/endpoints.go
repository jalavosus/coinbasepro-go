package endpoints

import (
	"github.com/jalavosus/coinbasepro-go/util/secrets"
)

const (
	productionRestAPIURI       string = "https://api.pro.coinbase.com"
	productionWebsocketFeedURI string = "wss://ws-feed.pro.coinbase.com"
)

const (
	sandboxRestAPIURI       string = "https://api-public.sandbox.pro.coinbase.com"
	sandboxWebsocketFeedURI string = "wss://ws-feed-public.sandbox.pro.coinbase.com"
)

// RestAPIURI returns a Coinbase Pro REST API URI based on whether or not the COINBASE_USE_SANDBOX
// environment variable is set to true or not (or is unset)
func RestAPIURI() string {
	if secrets.UseSandbox() {
		return sandboxRestAPIURI
	}

	return productionRestAPIURI
}

// WebsocketFeedURI returns a Coinbase Pro Websocket Feed API based on whether or not the COINBASE_USE_SANDBOX
// environment variable is set to true or not (or is unset)
func WebsocketFeedURI() string {
	if secrets.UseSandbox() {
		return sandboxWebsocketFeedURI
	}

	return productionWebsocketFeedURI
}

// Endpoints for account interaction
const (
	Accounts       string = "/accounts"
	Account               = Accounts + "/%[1]s" // formatted value must be an account ID.
	AccountHistory        = Account + "/ledger"
	AccountHolds          = Account + "/holds"
)

// Endpoints for order interaction
const (
	Orders string = "/orders"
	Order         = Orders + "/%[1]s" // formatted value must be an order ID.
)

// Endpoints for fills interaction
const (
	Fills string = "/fills"
)

// Endpoints for transfers interaction
const (
	Transfers string = "/transfers"
)

//nolint:unused,varcheck
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

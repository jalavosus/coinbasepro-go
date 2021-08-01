package coinbasepro

import (
	"time"

	"github.com/shopspring/decimal"
)

// Account represents a Coinbase Pro portfolio account for a given currency.
// This model is based off of the example response data from https://docs.pro.coinbase.com/#accounts.
type Account struct {
	Id             string          `json:"id"`
	Currency       string          `json:"currency"`
	Balance        decimal.Decimal `json:"balance"`
	Available      decimal.Decimal `json:"available"`
	Hold           decimal.Decimal `json:"hold"`
	ProfileId      string          `json:"profile_id"`
	TradingEnabled bool            `json:"trading_enabled"`
}

// Product represents a Coinbase Pro product object.
// This model is based off of the example response data from https://docs.pro.coinbase.com/#products.
type Product struct {
	Id              string          `json:"id"`
	DisplayName     string          `json:"display_name"`
	BaseCurrency    string          `json:"base_currency"`
	QuoteCurrency   string          `json:"quote_currency"`
	BaseIncrement   decimal.Decimal `json:"base_increment"`
	QuoteIncrement  decimal.Decimal `json:"quote_increment"`
	BaseMinSize     decimal.Decimal `json:"base_min_size"`
	BaseMaxSize     decimal.Decimal `json:"base_max_size"`
	MinMarketFunds  decimal.Decimal `json:"min_market_funds"`
	MaxMarketFunds  decimal.Decimal `json:"max_market_funds"`
	Status          string          `json:"status"`
	StatusMessage   string          `json:"status_message"`
	CancelOnly      bool            `json:"cancel_only"`
	LimitOnly       bool            `json:"limit_only"`
	PostOnly        bool            `json:"post_only"`
	TradingDisabled bool            `json:"trading_disabled"`
	FxStablecoin    bool            `json:"fx_stablecoin"`
}

type AccountHistoryEntry struct {
	Id        string                      `json:"id"`
	CreatedAt time.Time                   `json:"created_at"`
	Amount    decimal.Decimal             `json:"amount"`
	Balance   decimal.Decimal             `json:"balance"`
	Type      string                      `json:"type"`
	Details   *AccountHistoryEntryDetails `json:"details"`
}

type AccountHistoryEntryDetails struct {
	OrderId   string `json:"order_id"`
	TradeId   string `json:"trade_id"`
	ProductId string `json:"product_id"`
}

type AccountHold struct {
	Id        string          `json:"id"`
	AccountId string          `json:"account_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Amount    decimal.Decimal `json:"amount"`
	Type      string          `json:"type"`
	Ref       string          `json:"ref"`
}

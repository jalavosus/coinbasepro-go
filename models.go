package coinbasepro

import (
	"encoding/json"

	"github.com/shopspring/decimal"

	"github.com/jalavosus/coinbasepro-go/util"
)

// Account represents a Coinbase Pro portfolio account for a given currency.
// This model is based off of the example response data from https://docs.pro.coinbase.com/#accounts.
type Account struct {
	ID             string
	Currency       string
	Balance        decimal.Decimal
	Available      decimal.Decimal
	Hold           decimal.Decimal
	ProfileID      string
	TradingEnabled bool
}

func (a Account) MarshalJSON() ([]byte, error) {
	rawMap := map[string]interface{}{
		"id":              a.ID,
		"currency":        a.Currency,
		"balance":         a.Balance.String(),
		"available":       a.Available.String(),
		"hold":            a.Hold.String(),
		"profile_id":      a.ProfileID,
		"trading_enabled": a.TradingEnabled,
	}

	return json.Marshal(rawMap)
}

// UnmarshalJSON is specifically implemented to enable parsing fields
// which are returned by the API as strings into decimal.Decimal objects.
func (a *Account) UnmarshalJSON(data []byte) error {
	var rawMap map[string]interface{}

	if err := json.Unmarshal(data, &rawMap); err != nil {
		return err
	}

	a.ID = rawMap["id"].(string)
	a.Currency = rawMap["currency"].(string)
	a.Balance = util.StringToDecimal(rawMap["balance"].(string))
	a.Available = util.StringToDecimal(rawMap["available"].(string))
	a.Hold = util.StringToDecimal(rawMap["hold"].(string))
	a.TradingEnabled = rawMap["trading_enabled"].(bool)

	return nil
}

// Product represents a Coinbase Pro product object.
// This model is based off of the example response data from https://docs.pro.coinbase.com/#products.
type Product struct {
	ID              string
	DisplayName     string
	BaseCurrency    string
	QuoteCurrency   string
	BaseIncrement   decimal.Decimal
	QuoteIncrement  decimal.Decimal
	BaseMinSize     decimal.Decimal
	BaseMaxSize     decimal.Decimal
	MinMarketFunds  decimal.Decimal
	MaxMarketFunds  decimal.Decimal
	Status          string
	StatusMessage   string
	CancelOnly      bool
	LimitOnly       bool
	PostOnly        bool
	TradingDisabled bool
	FxStablecoin    bool
}

func (p Product) MarshalJSON() ([]byte, error) {
	rawMap := map[string]interface{}{
		"id":               p.ID,
		"display_name":     p.DisplayName,
		"base_currency":    p.BaseCurrency,
		"quote_currency":   p.QuoteCurrency,
		"base_increment":   p.BaseIncrement.String(),
		"quote_increment":  p.QuoteIncrement.String(),
		"base_min_size":    p.BaseMinSize.String(),
		"base_max_size":    p.BaseMaxSize.String(),
		"min_market_funds": p.MinMarketFunds.String(),
		"max_market_funds": p.MaxMarketFunds.String(),
		"status":           p.Status,
		"status_message":   p.StatusMessage,
		"cancel_only":      p.CancelOnly,
		"limit_only":       p.LimitOnly,
		"post_only":        p.PostOnly,
		"trading_disabled": p.TradingDisabled,
		"fx_stablecoin":    p.FxStablecoin,
	}

	return json.Marshal(rawMap)
}

// UnmarshalJSON is specifically implemented to enable parsing fields
// which are returned by the API as strings into decimal.Decimal objects.
func (p *Product) UnmarshalJSON(data []byte) error {
	var rawMap map[string]interface{}

	if err := json.Unmarshal(data, &rawMap); err != nil {
		return err
	}

	p.ID = rawMap["id"].(string)
	p.DisplayName = rawMap["display_name"].(string)
	p.BaseCurrency = rawMap["base_currency"].(string)
	p.QuoteCurrency = rawMap["quote_currency"].(string)
	p.BaseIncrement = util.StringToDecimal(rawMap["base_increment"].(string))
	p.QuoteIncrement = util.StringToDecimal(rawMap["quote_increment"].(string))
	p.BaseMinSize = util.StringToDecimal(rawMap["base_min_size"].(string))
	p.BaseMaxSize = util.StringToDecimal(rawMap["base_max_size"].(string))
	p.MinMarketFunds = util.StringToDecimal(rawMap["min_market_funds"].(string))
	p.MaxMarketFunds = util.StringToDecimal(rawMap["max_market_funds"].(string))
	p.Status = rawMap["status"].(string)
	p.StatusMessage = rawMap["status_message"].(string)
	p.CancelOnly = rawMap["cancel_only"].(bool)
	p.LimitOnly = rawMap["limit_only"].(bool)
	p.PostOnly = rawMap["post_only"].(bool)
	p.TradingDisabled = rawMap["trading_disabled"].(bool)
	p.FxStablecoin = rawMap["fx_stablecoin"].(bool)

	return nil
}

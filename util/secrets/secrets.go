// Package secrets contains functions for loading various secrets from
// the environment.
package secrets

import (
	"os"

	"github.com/pkg/errors"
)

const (
	envCoinbaseAPIKey        string = "COINBASE_API_KEY"
	envCoinbaseAPISecret     string = "COINBASE_API_SECRET"
	envCoinbaseAPIPassphrase string = "COINBASE_API_PASSPHRASE"
	envCoinbasePortfolioID   string = "COINBASE_PORTFOLIO_ID"
	envUseSandbox            string = "COINBASE_USE_SANDBOX"
)

var (
	coinbaseAPIKey        string
	coinbaseAPISecret     string
	coinbaseAPIPassphrase string
	coinbasePortfolioID   string
	useSandbox            *bool
)

// CoinbaseAPIKey returns a Coinbase Pro API key,
// loading it from the environment if necessary.
func CoinbaseAPIKey() string {
	if coinbaseAPIKey == "" {
		coinbaseAPIKey = loadFromEnv(envCoinbaseAPIKey)
	}

	return coinbaseAPIKey
}

// CoinbaseAPISecret returns a Coinbase Pro API secret,
// loading it from the environment if necessary.
func CoinbaseAPISecret() string {
	if coinbaseAPISecret == "" {
		coinbaseAPISecret = loadFromEnv(envCoinbaseAPISecret)
	}

	return coinbaseAPISecret
}

// CoinbaseAPIPassphrase returns a Coinbase Pro API passphrase,
// loading it from the environment if necessary.
func CoinbaseAPIPassphrase() string {
	if coinbaseAPIPassphrase == "" {
		coinbaseAPIPassphrase = loadFromEnv(envCoinbaseAPIPassphrase)
	}

	return coinbaseAPIPassphrase
}

// CoinbasePortfolioID returns a Coinbase Pro Account ID,
// loading it from the environment if necessary.
func CoinbasePortfolioID() string {
	if coinbasePortfolioID == "" {
		coinbasePortfolioID = loadFromEnv(envCoinbasePortfolioID)
	}

	return coinbasePortfolioID
}

func UseSandbox() bool {
	if useSandbox == nil {
		var val bool
		raw := loadFromEnv(envUseSandbox)
		switch raw {
		case "true", "TRUE", "yes", "1":
			val = true
		case "false", "FALSE", "no", "0":
			val = false
		}

		useSandbox = &val
	}

	return *useSandbox
}

func loadFromEnv(envKey string) string {
	res := os.Getenv(envKey)
	if res == "" {
		panic(errors.Errorf("%[1]s not found in environment. Is it set?", envKey))
	}

	return res
}

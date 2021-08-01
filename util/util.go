package util

import (
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// CoinbaseTimeFormat taken from the Coinbase Pro API reference:
// https://docs.pro.coinbase.com/#timestamps
const CoinbaseTimeFormat string = "2006-01-02T15:04:05.000Z"

// StringToDecimal parses a string into a decimal.Decimal,
// panicking if it fails.
func StringToDecimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		panic(errors.Wrapf(err, "error parsing %[1]s to decimal.Decimal", s))
	}

	return d
}

// ParseTimestampFromString takes a timestamp string from
// a Coinbase Pro API response and returns a time.Time object.
func ParseTimestampFromString(s string) time.Time {
	t, err := time.Parse(CoinbaseTimeFormat, s)
	if err != nil {
		panic(errors.Wrapf(err, "unable to parse time string %[1]s to time.Time", s))
	}

	return t
}

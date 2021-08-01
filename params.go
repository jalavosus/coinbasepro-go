package coinbasepro

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jalavosus/coinbasepro-go/util"
)

type GetAccountHistoryParams struct {
	// Before requires a positive integer. If set, returns ledger entries before the specified integer.
	Before *uint
	// After requires a positive integer. If set, returns ledger entries after the specified integer.
	After *uint
	// If set, returns ledger entries created after the start_date timestamp, sorted by newest creation date. When combined with end_date, returns ledger entries in the specified time range.
	StartDate *time.Time
	// If set, returns ledger entries created before the end_date timestamp, sorted by newest creation date.
	EndDate *time.Time
	// Limit defaults to 1000
	Limit *uint
}

func (p GetAccountHistoryParams) GetEncoded() string {
	params := url.Values{}
	if p.Before != nil {
		params.Set("before", fmt.Sprintf("%d", *p.Before))
	}
	if p.After != nil {
		params.Set("after", fmt.Sprintf("%d", *p.After))
	}
	if p.StartDate != nil {
		params.Set("start_date", util.TimeToTimestampString(*p.StartDate))
	}
	if p.EndDate != nil {
		params.Set("end_date", util.TimeToTimestampString(*p.EndDate))
	}
	if p.Limit != nil && *p.Limit != 0 {
		params.Set("limit", fmt.Sprintf("%d", *p.Limit))
	}

	return params.Encode()
}

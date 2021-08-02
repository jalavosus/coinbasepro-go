package coinbasepro

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/jalavosus/coinbasepro-go/internal/apirequests"
	"github.com/jalavosus/coinbasepro-go/internal/endpoints"
)

// GetAccounts wraps the /accounts API endpoint,
// returning all accounts which the loaded API key
// has access to.
func GetAccounts() ([]*Account, error) {
	var accounts []*Account

	respBody, err := apirequests.GETRequest(endpoints.Accounts, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetAccounts() error")
	}

	if err := json.Unmarshal(respBody, &accounts); err != nil {
		return nil, errors.Wrap(err, "GetAccounts() error")
	}

	return accounts, nil
}

// GetAccount wraps the /accounts/<account-id> API endpoint,
// returning an Account representing the portfolio account
// with the passed accountID.
func GetAccount(accountID string) (*Account, error) {
	var account *Account

	endpoint := fmt.Sprintf(endpoints.Account, accountID)

	respBody, err := apirequests.GETRequest(endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetAccount() error")
	}

	if err := json.Unmarshal(respBody, &account); err != nil {
		return nil, errors.Wrap(err, "GetAccount() error")
	}

	return account, nil
}

// GetAccountFromCurrency attempts to locate an account in the user's portfolio
// based on its currency.
// Note that this function can return a nil Account if one with the passed
// currency isn't found.
func GetAccountFromCurrency(currency string) (*Account, error) {
	var account *Account

	allAccounts, err := GetAccounts()
	if err != nil {
		return nil, err
	}

	for _, a := range allAccounts {
		if a.Currency == currency {
			account = a

			break
		}
	}

	return account, nil
}

// GetAccountHistory wraps the /accounts/<account-id>/ledger API endpoint,
// returning an array of ledger entries for the specified account.
func GetAccountHistory(accountID string, params *GetAccountHistoryParams) ([]*AccountHistoryEntry, error) {
	var historyEntries []*AccountHistoryEntry

	endpoint := fmt.Sprintf(endpoints.AccountHistory, accountID)

	respBody, err := apirequests.GETRequest(endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetAccountHistory() error")
	}

	if err := json.Unmarshal(respBody, &historyEntries); err != nil {
		return nil, errors.Wrap(err, "GetAccountHistory() error")
	}

	return historyEntries, nil
}

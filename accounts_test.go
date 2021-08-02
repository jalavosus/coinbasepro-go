package coinbasepro_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/coinbasepro-go"
)

var (
	btcAccountID     string
	ethAccountID     string
	invalidAccountID string
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	btcAccountID = os.Getenv("TEST_BTC_ACCOUNT_ID")
	ethAccountID = os.Getenv("TEST_ETH_ACCOUNT_ID")
	invalidAccountID = os.Getenv("TEST_INVALID_ACCOUNT_ID")

	code := m.Run()

	os.Exit(code)
}

func TestGetAccount(t *testing.T) {
	type args struct {
		accountID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "BTC",
			args:    args{btcAccountID},
			wantErr: false,
		},
		{
			name:    "ETH",
			args:    args{ethAccountID},
			wantErr: false,
		},
		{
			name:    "INVALID",
			args:    args{invalidAccountID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := coinbasepro.GetAccount(tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetAccountFromCurrency(t *testing.T) {
	type args struct {
		currency string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantNil bool
	}{
		{
			name:    "ETH",
			args:    args{"ETH"},
			wantErr: false,
			wantNil: false,
		},
		{
			name:    "INVALID",
			args:    args{"INVALID"},
			wantErr: false,
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := coinbasepro.GetAccountFromCurrency(tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountFromCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("GetAccountFromCurrency() got == %v, wantNil %v", got, tt.wantNil)
			}
		})
	}
}

func TestGetAccounts(t *testing.T) {
	got, err := coinbasepro.GetAccounts()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.NotEmpty(t, got)
}

func TestGetAccountHistory(t *testing.T) {
	type args struct {
		accountID string
		params    *coinbasepro.GetAccountHistoryParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "BTC",
			args:    args{accountID: btcAccountID},
			wantErr: false,
		},
		{
			name:    "ETH",
			args:    args{accountID: ethAccountID},
			wantErr: false,
		},
		{
			name:    "INVALID",
			args:    args{accountID: invalidAccountID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := coinbasepro.GetAccountHistory(tt.args.accountID, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

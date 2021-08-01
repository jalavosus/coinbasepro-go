package coinbasepro_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/coinbasepro-go"
)

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
			name:    "USDC",
			args:    args{os.Getenv("TEST_USDC_ACCOUNT_ID")},
			wantErr: false,
		},
		{
			name:    "SHIB",
			args:    args{os.Getenv("TEST_SHIB_ACCOUNT_ID")},
			wantErr: false,
		},
		{
			name:    "INVALID",
			args:    args{os.Getenv("TEST_INVALID_ACCOUNT_ID")},
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
			name:    "USDC",
			args:    args{"USDC"},
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

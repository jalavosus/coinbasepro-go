package coinbasepro_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/coinbasepro-go"
)

func TestGetProducts(t *testing.T) {
	got, err := coinbasepro.GetProducts()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.NotEmpty(t, got)
}

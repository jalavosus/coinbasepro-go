package coinbasepro

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/jalavosus/coinbasepro-go/internal/apirequests"
	"github.com/jalavosus/coinbasepro-go/internal/endpoints"
)

// GetProducts wraps the /products API endpoint, returning an array of Coinbase Pro
// Product objects.
func GetProducts() ([]*Product, error) {
	var products []*Product

	respBody, err := apirequests.GETRequest(endpoints.Products, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetProducts() error")
	}

	if err := json.Unmarshal(respBody, &products); err != nil {
		return nil, errors.Wrap(err, "GetProducts() error")
	}

	return products, nil
}

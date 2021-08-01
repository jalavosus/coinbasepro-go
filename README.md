# coinbasepro-go

Go wrapper around the Coinbase Pro API, using [fasthttp](https://github.com/valyala/fasthttp) for 
firing off API requests.

Note that this library is a continuing work-in-progress and is provided AS-IS. I take no responsibility for any fiat money or 
cryptocurrency lost as a result of a malfunction or error in this library.

## Installation
`go get -u github.com/jalavosus/coinbasepro-go`

## Usage

The following keys must be set in the environment, either with a `.env` file or via another method:

- `COINBASE_API_KEY`: A Coinbase Pro API key
- `COINBASE_API_SECRET`: The alphanumeric secret key provided by Coinbase Pro upon creation of an API key.
- `COINBASE_API_PASSPHRASE`: The passphrase assigned by the API key's creator upon creation. 
- `COINBASE_PORTFOLIO_ID`: ID of the portfolio the provided API key is allowed to access.

Assuming those keys are set in the environment, usage is as simple as 

```go
package main

import (
	"fmt"
	
	"github.com/jalavosus/coinbasepro-go"
)

func main() {
	allAccounts := coinbasepro.GetAccounts()
	
	for _, account := range allAccounts {
		fmt.Println(account.Currency)
	}
}
```
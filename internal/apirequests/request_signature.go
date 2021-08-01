package apirequests

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/jalavosus/coinbasepro-go/util/secrets"
)

// BuildRequestSignature builds the signature required by the Coinbase Pro API for the CB-ACCESS-SIGN
// header, using the request's method, path, and JSON body (if provided).
// The generated signature, as well as the timestamp (in seconds) used for the signature (for use with the
// CB-ACCESS-TIMESTAMP header) are returned.
func BuildRequestSignature(requestMethod, requestPath string, requestBody []byte) (string, string) {
	secret, err := base64.StdEncoding.DecodeString(secrets.CoinbaseAPISecret())
	if err != nil {
		panic(errors.Wrap(err, "base64 decoding error"))
	}

	timestamp := time.Now().Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)

	prehashStr := timestampStr + strings.ToUpper(requestMethod) + requestPath
	if requestBody != nil {
		prehashStr += string(requestBody)
	}

	prehashBytes := []byte(prehashStr)

	mac := hmac.New(sha256.New, secret)
	_, err = mac.Write(prehashBytes)
	if err != nil {
		panic(errors.Wrap(err, "error attempting to write prehash bytes to hmac"))
	}

	sigBytes := mac.Sum(nil)

	sigEncoded := base64.StdEncoding.EncodeToString(sigBytes)

	return sigEncoded, timestampStr
}

package apirequests

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/jalavosus/coinbasepro-go/internal/endpoints"
	"github.com/jalavosus/coinbasepro-go/internal/headers"
	"github.com/jalavosus/coinbasepro-go/util/secrets"
)

func GETRequest(endpoint string, params RequestParams) ([]byte, error) {
	return request(endpoint, "GET", params, nil)
}

// POSTRequest fires off a POST request to the specifed endpoint. If body is non-nil,
// it is added to the request.
func POSTRequest(endpoint string, body interface{}, params RequestParams) ([]byte, error) {
	var reqBody []byte = nil
	if body != nil {
		var marshalErr error
		reqBody, marshalErr = json.Marshal(body)
		if marshalErr != nil {
			return nil, errors.Wrap(marshalErr, "POSTRequest(): error marshaling request body")
		}
	}

	return request(endpoint, "POST", params, reqBody)

}

// DELETERequest fires off a DELETE request to the specified endpoint.
func DELETERequest(endpoint string, params RequestParams) ([]byte, error) {
	return request(endpoint, "DELETE", params, nil)
}

// request does all of the hard work for GETRequest, POSTRequest, etc.
func request(endpoint, method string, params RequestParams, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	uri := buildURIWithParams(endpoints.RestAPI+endpoint, params)
	req.SetRequestURI(uri.String())
	setHeaders(uri.Path, strings.ToUpper(method), body, req)

	if body != nil {
		req.SetBody(body)
	}

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, errors.Wrap(err, "fasthttp.Do error")
	}

	respBody := resp.Body()

	hasErr, errMsg := checkResponseWithMessage(respBody)
	if hasErr {
		return nil, errors.New(errMsg)
	}

	return respBody, nil
}

func setHeaders(basePath, method string, body []byte, req *fasthttp.Request) {
	req.Header.Set(headers.AccessKey, secrets.CoinbaseAPIKey())
	req.Header.Set(headers.AccessPassphrase, secrets.CoinbaseAPIPassphrase())

	signature, timestamp := BuildRequestSignature(method, basePath, body)
	req.Header.Set(headers.AccessSignature, signature)
	req.Header.Set(headers.AccessTimestamp, timestamp)
}

func buildURIWithParams(endpoint string, params RequestParams) *url.URL {
	var rawUri = endpoint

	if params != nil {
		encodedParams := params.GetEncoded()
		if len(encodedParams) > 0 {
			rawUri += "?" + encodedParams
		}

	}

	uri, err := url.Parse(rawUri)
	if err != nil {
		panic(errors.Wrap(err, "Unable to build parsed URI"))
	}

	return uri
}

// checkResponseWithMessage checks an API response for a "message" field,
// and returns that field if found.
func checkResponseWithMessage(respBody []byte) (bool, string) {
	var resp ResponseWithMessage

	// Screw it, who cares.
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return false, ""
	}

	return resp.Message != "", resp.Message
}

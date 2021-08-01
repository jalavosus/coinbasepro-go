package apirequests

type RequestParams interface {
	GetEncoded() string
}

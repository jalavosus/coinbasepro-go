package headers

const coinbaseHeaderPrefix string = "CB-"

const (
	accessPrefix     = coinbaseHeaderPrefix + "ACCESS-"
	AccessKey        = accessPrefix + "KEY"
	AccessSignature  = accessPrefix + "SIGN"
	AccessTimestamp  = accessPrefix + "TIMESTAMP"
	AccessPassphrase = accessPrefix + "PASSPHRASE"
)

const (
	PaginationBefore = coinbaseHeaderPrefix + "BEFORE"
	PaginationAfter  = coinbaseHeaderPrefix + "AFTER"
)

package postgres

type ctxKey string

const (
	txKey         ctxKey = "tx_key"
	txCommitedKey ctxKey = "tx_commited_key"
)

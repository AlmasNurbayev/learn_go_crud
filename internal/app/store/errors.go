package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrScan           = errors.New("error if scan rows")
)

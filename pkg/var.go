package pkg

import "errors"

const (
	ResultFile  = "result.json"
	CopyFileDir = "copy"
)

var (
	ErrExistDuplicateFile = errors.New("exist duplicate file")
)

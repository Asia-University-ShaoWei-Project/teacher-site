package mock

import (
	"context"
	"time"
)

var Ctx = context.Background()

const (
	Date      = "2020-01-01"
	EmptyStr  = ""
	WordStr   = "word"
	EmptyJson = "{}"

	NumPk         = 1
	PkZeroStr     = "0"
	PkStr         = "1"
	NegativePkStr = "-1"
	UnknownNumPK  = 999
)

// message
const (
	UpdateMsg = "update message"
	FailMsg   = "fail message"
	CreateMsg = "create message"
)

func NewMsg() string {
	format := `03:04:05PM`
	t := time.Now().Format(format)
	msg := `new message at ` + t
	return msg
}

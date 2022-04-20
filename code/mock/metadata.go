package mock

import "time"

const (
	Date      = "2020-01-01"
	EmptyStr  = ""
	StrWord   = "a"
	EmptyJson = "{}"

	NumPK        = 1
	UnknownNumPK = 999
)

// message
const (
	UpdateMsg = "update message"
	FailMsg   = "fail message"
	CreateMsg = "create message"
)

func NewMsg() string {
	format := `01/02 03:04:05PM`
	t := time.Now().Format(format)
	msg := `new message at ` + t
	return msg
}

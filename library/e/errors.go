package e

import "fmt"

type Error struct {
	Msg string
	ErrCode int
}

func New(message string, code int) error {
	return Error{
		Msg: message,
		ErrCode: code,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s (错误码: %d)", e.Msg, e.ErrCode)
}
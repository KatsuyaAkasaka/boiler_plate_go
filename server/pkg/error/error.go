package e

import (
	"strings"
)

type ErrStruct struct {
	HttpStatusCode int
	ErrCode        int
	Description    string
}

type Err *ErrStruct

func createErr(status int, msg int, descs ...string) Err {
	return &ErrStruct{
		HttpStatusCode: status,
		ErrCode:        msg,
		Description:    strings.Join(descs, "."),
	}
}

func GetError(e Err) ErrStruct {
	return *e
}

func (e ErrStruct) Error() int {
	return e.ErrCode
}

var Stripe = struct {
}{}

package e

import "fmt"

// type NewTemplate interface {
// 	NotFound() Err
// 	Duplicated() Err
// 	InvalidParameter() Err
// }

type ErrTemplate struct {
	NotFound         Err
	Duplicated       Err
	InvalidParameter Err
}

type Prefix int

func (p Prefix) NotFound() Err {
	return New(404, int(p)+1)
}

func (p Prefix) Duplicated() Err {
	return New(409, int(p)+2)
}

func (p Prefix) InvalidParameter() Err {
	return New(400, int(p)+3)
}

const (
	dbPrefix = Prefix(iota * 100)
	sqlPrefix
	systemPrefix
	userPrefix
)

type ErrMessage struct {
	HttpStatusCode int
	ErrCode        int
	// ErrMetadata
}

func (e *ErrMessage) Error() string {
	return fmt.Sprintln(e.ErrCode)
}

func GetErr(e Err) error {
	return fmt.Errorf("errCode: %d", e.ErrCode)
}

type Err *ErrMessage

func New(status int, errCode int) Err {
	return &ErrMessage{
		HttpStatusCode: status,
		ErrCode:        errCode,
	}
}

func newTemplate(p Prefix) ErrTemplate {
	return ErrTemplate{
		NotFound:         p.NotFound(),
		Duplicated:       p.Duplicated(),
		InvalidParameter: p.InvalidParameter(),
	}
}

// func createMetadata(prefix int, name string) ErrMetadata {
// 	return ErrMetadata{
// 		Prefix: prefix,
// 		Name:   name,
// 	}
// }

// func New(text string) error {
// 	return &errorString{text}
// }

// // errorString is a trivial implementation of error.
// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

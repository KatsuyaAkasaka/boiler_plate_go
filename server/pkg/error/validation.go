package e

const validationPrefix = 200

const (
	BadRequest = iota + validationPrefix
)

type validationError struct {
	BadRequest Err
}

var (
	Validation = validationError{
		BadRequest: createErr(400, BadRequest),
	}
)

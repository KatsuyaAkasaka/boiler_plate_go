package e

const (
	BadRequest = int(iota + systemPrefix)
	BadResponse
	Unexported
	TokenInvaid
	LoginRequired
	SessionExpired

	S3Err
	StripeErr

	FailedToSendMail
)

type SystemErr struct {
	BadRequest     Err
	BadResponse    Err
	Unexported     Err
	TokenInvalid   Err
	LoginRequired  Err
	SessionExpired Err

	S3Err     Err
	StripeErr Err

	FailedToSendMail Err
}

var System = SystemErr{
	New(400, BadRequest),
	New(400, BadRequest),
	New(500, Unexported),
	New(401, TokenInvaid),
	New(401, LoginRequired),
	New(401, SessionExpired),

	New(500, S3Err),
	New(500, StripeErr),

	New(500, FailedToSendMail),
}

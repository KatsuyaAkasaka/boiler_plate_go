package e

const sqlPrefix = 100

const (
	FailedToExecute = iota + sqlPrefix
	FailedToBeginTx
	FailedToCommitTx
)

type sqlError struct {
	FailedToExecute  Err
	FailedToBeginTx  Err
	FailedToCommitTx Err
}

var (
	SQL = sqlError{
		FailedToExecute:  createErr(500, FailedToExecute),
		FailedToBeginTx:  createErr(500, FailedToBeginTx),
		FailedToCommitTx: createErr(500, FailedToCommitTx),
	}
)

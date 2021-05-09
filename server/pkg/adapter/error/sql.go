package e

const (
	FailedToExecute = int(iota + sqlPrefix)
	FailedToBeginTx
	FailedToCommitTx
)

type SQLErr struct {
	FailedToExecute  Err
	FailedToBeginTx  Err
	FailedToCommitTx Err
}

var SQL = SQLErr{
	New(500, FailedToExecute),
	New(500, FailedToBeginTx),
	New(500, FailedToCommitTx),
}

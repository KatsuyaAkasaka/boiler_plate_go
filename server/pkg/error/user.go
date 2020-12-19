package e

const UserPrefix = 400

const (
	failedToCreateUser = iota + UserPrefix
	failedToUpdateUser
	failedToDeleteUser
	failedToFindUser
)

type userErr struct {
	FailedToCreate Err
	FailedToUpdate Err
	FailedToDelete Err
	FailedToFind   Err
}

var (
	User = userErr{
		FailedToCreate: createErr(500, failedToCreateUser),
		FailedToUpdate: createErr(500, failedToUpdateUser),
		FailedToDelete: createErr(500, failedToDeleteUser),
		FailedToFind:   createErr(500, failedToFindUser),
	}
)

func (u userErr) CheckDBError(err error) Err {
	return checkDBError(err, UserPrefix)
}

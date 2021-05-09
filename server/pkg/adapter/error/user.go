package e

type UserErr struct {
	ErrTemplate
}

var User = UserErr{
	newTemplate(userPrefix),
}

func (u UserErr) CheckDBError(err error, includesNotFound bool) Err {
	return checkDBError(err, includesNotFound, userPrefix)
}

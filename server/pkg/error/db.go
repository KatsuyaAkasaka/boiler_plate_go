package e

import (
	"strings"

	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/logger"
	"github.com/jinzhu/gorm"
)

const (
	RecordNotFound = iota
	InvalidSQL
	CantStartTransaction
	InvalidTransaction
	Unaddressable
	DuplicatedKey
	LockTimeout
	UnexpectedDBError
)

func DBError(errType int, prefix int) Err {
	errCode := errType + prefix
	switch errType {
	case RecordNotFound:
		return createErr(404, errCode)
	case InvalidSQL:
		return createErr(400, errCode)
	case CantStartTransaction:
		return createErr(400, errCode)
	case InvalidTransaction:
		return createErr(400, errCode)
	case Unaddressable:
		return createErr(400, errCode)
	case DuplicatedKey:
		return createErr(409, errCode)
	case LockTimeout:
		return createErr(408, errCode)
	case UnexpectedDBError:
		return createErr(500, errCode)
	default:
		return createErr(500, errCode)
	}
}

func checkDBError(err error, p int) Err {
	switch err {
	case gorm.ErrRecordNotFound:
		log.Warn(err.Error())
		return DBError(RecordNotFound, p)
	case gorm.ErrInvalidSQL:
		log.Error(err.Error())
		return DBError(InvalidSQL, p)
	case gorm.ErrCantStartTransaction:
		log.Error(err.Error())
		return DBError(CantStartTransaction, p)
	case gorm.ErrInvalidTransaction:
		log.Error(err.Error())
		return DBError(InvalidTransaction, p)
	case gorm.ErrUnaddressable:
		log.Error(err.Error())
		return DBError(Unaddressable, p)
	}

	//重複エラーのハンドリング
	if strings.Contains(err.Error(), "1062") {
		log.Warn(err.Error())
		return DBError(DuplicatedKey, p)
	}
	// ロック待機のタイムアウト
	if strings.Contains(err.Error(), "1205") {
		log.Warn(err.Error())
		return DBError(LockTimeout, p)
	}

	log.Error(err.Error())
	return DBError(UnexpectedDBError, p)
}

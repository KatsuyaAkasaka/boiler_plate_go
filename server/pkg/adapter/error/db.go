package e

import (
	"strings"

	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"gorm.io/gorm"
)

const (
	RecordNotFound = int(iota + dbPrefix)
	InvalidTransaction
	DuplicatedKey
	LockTimeout
	UnexpectedDBError
)

func DBError(errCode int) Err {
	switch errCode {
	case RecordNotFound:
		return New(404, errCode)
	case InvalidTransaction:
		return New(400, errCode)
	case DuplicatedKey:
		return New(409, errCode)
	case LockTimeout:
		return New(408, errCode)
	case UnexpectedDBError:
		return New(500, errCode)
	default:
		return New(500, errCode)
	}
}

func checkDBError(err error, includesNotFound bool, p Prefix) Err {
	switch err {
	case gorm.ErrRecordNotFound:
		if includesNotFound {
			log.Warn(err.Error())
			return p.NotFound()
		}
		return nil

	case gorm.ErrInvalidTransaction:
		log.Error(err.Error())
		return DBError(InvalidTransaction)

	default:
		//重複エラーのハンドリング
		if strings.Contains(err.Error(), "1062") {
			log.Warn(err.Error())
			return p.Duplicated()
		}
		// ロック待機のタイムアウト
		if strings.Contains(err.Error(), "1205") {
			log.Warn(err.Error())
			return DBError(LockTimeout)
		}

		log.Error(err.Error())
		return DBError(UnexpectedDBError)
	}
}

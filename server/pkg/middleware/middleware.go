package middleware

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
)

type Middles struct {
	JWT *JWT
}

func NewMiddleware(repos *repository.Repositories) *Middles {
	return &Middles{
		JWT: NewJWT(repos),
	}
}

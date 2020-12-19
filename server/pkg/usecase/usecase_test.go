package usecase

import (
	"os"
	"testing"

	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/logger"
	mock "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/mock"
	"github.com/golang/mock/gomock"
)

var (
	ctrl *gomock.Controller
)

func TestMain(m *testing.M) {
	log.NewLogger("test")
	code := m.Run()
	os.Exit(code)
}

func initRepo(t *testing.T) {
	ctrl = gomock.NewController(t)
	defer ctrl.Finish()
}

func newMockUserRepository() *mock.MockUserRepository {
	return mock.NewMockUserRepository(ctrl)
}

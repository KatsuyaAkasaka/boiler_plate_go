package entry

import (
	"os"
	"testing"

	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/datastore"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
	"gorm.io/gorm"
)

var testRepo *repository.Repositories = newTestRepository()
var testDB *gorm.DB

func TestMain(m *testing.M) {
	log.NewLogger("test")
	deleteAllTables()
	createAllTables()

	code := m.Run()

	os.Exit(code)
}

func newTestRepository() *repository.Repositories {
	confMap := config.GetConf("test")
	dbConf := confMap.DB

	dbInfo := datastore.DBInfo{
		UserName: dbConf["user"].(string),
		Password: dbConf["password"].(string),
		Host:     dbConf["host"].(string),
		Port:     dbConf["port"].(string),
		Name:     dbConf["name"].(string),
	}
	testDB = datastore.InitDB(&dbInfo)
	testDB.LogMode(false)
	return NewRepository(testDB)
}

var resetTables = struct {
	User *entity.User
}{
	&entity.User{},
}

func deleteAllTables() {
	testDB.DropTable(&entity.User{})
}

func createAllTables() {
	testDB.CreateTable(&entity.User{})
}

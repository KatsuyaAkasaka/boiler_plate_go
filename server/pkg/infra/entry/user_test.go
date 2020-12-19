package entry

import (
	"testing"

	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
	"github.com/go-playground/assert/v2"
)

func reset() {
	testDB.DropTable(&entity.User{})
	testDB.AutoMigrate(&entity.User{})
}

func TestCreate(t *testing.T) {
	testUser := &entity.User{
		UserID:          "sakas",
		Email:           "sakas@gmail.com",
		Gender:          1,
		Description:     "test",
		ProfileImageURI: "https://test/test.png",
	}
	t.Run("正常系", func(t *testing.T) {
		reset()
		user, err := testRepo.User.Create(testUser)

		assert.Equal(t, err, nil)
		assert.Equal(t, user, testUser)
		var insertedUser entity.User
		testDB.Where("user_id = ?", testUser.UserID).First(&insertedUser)
		// 成功
		assert.Equal(t, testUser.Email, insertedUser.Email)
		assert.Equal(t, testUser.Gender, insertedUser.Gender)
		assert.Equal(t, testUser.Description, insertedUser.Description)

		user, err = testRepo.User.Create(testUser)
		assert.Equal(t, err, e.DBError(e.DuplicatedKey, e.UserPrefix))
		assert.Equal(t, user, nil)
		return
	})
}

func TestUpdate(t *testing.T) {
	testUser := &entity.User{
		UserID:          "sakas",
		Email:           "sakas@gmail.com",
		Gender:          1,
		Description:     "test",
		ProfileImageURI: "https://test/test.png",
	}
	t.Run("正常系", func(t *testing.T) {
		reset()
		user, err := testRepo.User.Create(testUser)

		assert.Equal(t, err, nil)
		assert.Equal(t, user, testUser)
		var insertedUser entity.User
		testDB.Where("user_id = ?", testUser.UserID).First(&insertedUser)
		// 成功
		assert.Equal(t, testUser.Email, insertedUser.Email)
		assert.Equal(t, testUser.Gender, insertedUser.Gender)
		assert.Equal(t, testUser.Description, insertedUser.Description)

		user, err = testRepo.User.Create(testUser)
		assert.Equal(t, err, e.DBError(e.DuplicatedKey, e.UserPrefix))
		assert.Equal(t, user, nil)
		return
	})
}

func TestDelete(t *testing.T) {
	testUser := &entity.User{
		UserID:          "sakas",
		Email:           "sakas@gmail.com",
		Gender:          1,
		Description:     "test",
		ProfileImageURI: "https://test/test.png",
	}
	t.Run("正常系", func(t *testing.T) {
		reset()
		testDB.Create(testUser)
		createdUser := &entity.User{}
		testDB.Where("user_id = ?", testUser.UserID).First(createdUser)
		assert.Equal(t, createdUser.UserID, testUser.UserID)

		err := testRepo.User.Delete(entity.GetUserID(testUser.UserID))
		assert.Equal(t, err, nil)

		afterUser := &entity.User{}
		if err := testDB.Where("user_id = ?", testUser.UserID).First(afterUser).Error; err != nil {
			code := e.GetError(e.User.CheckDBError(err)).Error()
			assert.Equal(t, code, 400) // NotFound
		}
		return
	})
}

func TestFindByUserID(t *testing.T) {
	testUser := &entity.User{
		UserID:          "sakas",
		Email:           "sakas@gmail.com",
		Gender:          1,
		Description:     "test",
		ProfileImageURI: "https://test/test.png",
	}
	t.Run("正常系", func(t *testing.T) {
		reset()
		testDB.Create(testUser)
		user, err := testRepo.User.FindByUserID(entity.GetUserID(testUser.UserID))
		assert.Equal(t, err, nil)
		assert.Equal(t, user.UserID, testUser.UserID)
		return
	})
}

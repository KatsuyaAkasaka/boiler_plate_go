package usecase

import (
	"testing"

	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/input"
	"github.com/go-playground/assert/v2"
)

func TestGetUser(t *testing.T) {
	initRepo(t)
	mockRepo := newMockUserRepository()
	test := &entity.User{
		UserID:          "sakas",
		Email:           "test@gmail.com",
		NickName:        "testName",
		Gender:          1,
		Description:     "this is test",
		ProfileImageURI: "https://test/test.png",
	}

	mockRepo.EXPECT().FindByUserID(entity.UserID(test.UserID)).Return(test, nil)
	u := newUserUsecase(mockRepo)
	res, err := u.GetUser(test.UserID)
	assert.Equal(t, err, nil)
	assert.Equal(t, *res, *test)
}

func TestCreateUser(t *testing.T) {
	initRepo(t)
	mockRepo := newMockUserRepository()
	test := input.UserReq{
		UserID:          "sakas",
		Email:           "test@gmail.com",
		NickName:        "testName",
		Gender:          1,
		Description:     "this is test",
		ProfileImageURI: "https://test/test.png",
	}

	entityUser := test.GetUser()

	mockRepo.EXPECT().Create(entityUser).Return(entityUser, nil)
	u := newUserUsecase(mockRepo)
	res, err := u.CreateUser(&test)
	assert.Equal(t, err, nil)
	assert.Equal(t, res, entityUser)
}

func TestUpdateUser(t *testing.T) {
	initRepo(t)
	mockRepo := newMockUserRepository()
	test := input.UserReq{
		UserID:          "sakas",
		Email:           "test@gmail.com",
		NickName:        "testName",
		Gender:          1,
		Description:     "this is test",
		ProfileImageURI: "https://test/test.png",
	}

	entityUser := test.GetUser()

	mockRepo.EXPECT().Update(entityUser).Return(entityUser, nil)
	u := newUserUsecase(mockRepo)
	res, err := u.UpdateUser(&test)
	assert.Equal(t, err, nil)
	assert.Equal(t, res, entityUser)
}

func TestDeleUser(t *testing.T) {
	initRepo(t)
	mockRepo := newMockUserRepository()
	test := entity.User{
		UserID:          "sakas",
		Email:           "test@gmail.com",
		NickName:        "testName",
		Gender:          1,
		Description:     "this is test",
		ProfileImageURI: "https://test/test.png",
	}

	mockRepo.EXPECT().Delete(entity.UserID(test.UserID)).Return(nil)
	u := newUserUsecase(mockRepo)
	err := u.DeleteUser(test.UserID)
	assert.Equal(t, err, nil)
}

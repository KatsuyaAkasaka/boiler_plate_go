package handler

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/input"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	router      *gin.RouterGroup
	userUsecase *usecase.UserUsecase
}

func getUserRouter(r *gin.Engine, path string, ver string) *gin.RouterGroup {
	return r.Group("users")
}

func newUserHandler(r *gin.RouterGroup, uc *usecase.Usecases) {
	uh := &userHandler{
		router:      r,
		userUsecase: uc.User,
	}
	r.POST("", uh.createUser())
	r.GET("/:id", uh.getUser())
	r.PUT("/:id", uh.updateUser())
	r.DELETE("/:id", uh.deleteUser())
}

func (uh *userHandler) createUser() func(*gin.Context) {
	return func(c *gin.Context) {
		var user input.UserReq
		if err := c.BindJSON(&user); err != nil {
			returnErr(c, e.Validation.BadRequest)
			return
		}
		afterUser, err := uh.userUsecase.CreateUser(&user)
		if err != nil {
			returnErr(c, err)
			return
		}
		c.JSON(200, gin.H{
			"result": *afterUser,
		})
	}
}

func (uh *userHandler) getUser() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		afterUser, err := uh.userUsecase.GetUser(id)
		if err != nil {
			returnErr(c, err)
			return
		}
		c.JSON(200, gin.H{
			"result": *afterUser,
		})
	}
}

func (uh *userHandler) updateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user input.UserReq
		if err := c.BindJSON(&user); err != nil {
			returnErr(c, e.Validation.BadRequest)
			return
		}
		user.BindUserID(id)
		afterUser, err := uh.userUsecase.UpdateUser(&user)
		if err != nil {
			returnErr(c, err)
			return
		}
		c.JSON(200, gin.H{
			"result": *afterUser,
		})
	}
}

func (uh *userHandler) deleteUser() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := uh.userUsecase.DeleteUser(id)
		if err != nil {
			returnErr(c, err)
			return
		}
		returnSuccess(c)
	}
}

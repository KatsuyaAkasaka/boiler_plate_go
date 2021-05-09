package handler

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/middleware"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/input"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase   *usecase.UserUsecase
	systemUsecase *usecase.SystemUsecase
}

func newUserHandler(r *gin.RouterGroup, uc *usecase.Usecases, middles *middleware.Middles) {
	uh := &userHandler{
		userUsecase:   uc.User,
		systemUsecase: uc.System,
	}
	userRouter := r.Group("/users")
	// userRouterWithUUIDAuth := userRouter.Group(
	// 	"", middles.JWT.UUIDHandlerFunc(),
	// )
	userRouterWithEmailAuth := userRouter.Group(
		"", middles.JWT.EmailHandlerFunc(),
	)
	userRouterWithUUIDAuthNotErr := userRouter.Group(
		"", middles.JWT.UUIDHandlerNotErrFunc(),
	)
	// ユーザ作成 (これはjwtにemailが乗っているので注意)
	userRouterWithEmailAuth.POST("", uh.createUser())
	// 指定ユーザ情報取得
	userRouterWithUUIDAuthNotErr.GET("/:id", uh.getUser())
	// 指定アカウント削除
	userRouter.DELETE("/:id", uh.deleteUser())

	meRouter := r.Group(
		"/me", middles.JWT.UUIDHandlerFunc(),
	)
	// 自分のプロフィール取得
	meRouter.GET("/profile", uh.getMe())
	// 自分のプロフィール編集
	meRouter.PATCH("/profile", uh.updateMe())
	// メール通知の設定変更
	meRouter.PATCH("/email-notifications", uh.changeEmailNotify())
}

func (uh *userHandler) createUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, err := getEmail(c)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		var user input.UserReq
		if err := c.BindJSON(&user); err != nil {
			ReturnErr(c, e.System.BadRequest)
			return
		}
		user.BindEmail(email)
		res, err := uh.userUsecase.CreateUser(&user)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		// TODO: お触り会用にコミュニティも作成する
		// _, err = uh.communityUsecase.CreateCommunity((*entity.UUID)(&res.User.UUID), 2970)
		// if err != nil {
		// 	ReturnErr(c, err)
		// 	return
		// }
		ReturnSuccessData(c, &res)
	}
}

func (uh *userHandler) getMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := getUUID(c)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		res, err := uh.userUsecase.GetUserByUUID(uuid)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		ReturnSuccessData(c, &res)
	}
}

func (uh *userHandler) getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, err := entity.GetUserID(id)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		uuid, _ := getUUID(c)
		afterUser, err := uh.userUsecase.GetUserByUserID(uuid, userID)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		ReturnSuccessData(c, &afterUser)
	}
}

func (uh *userHandler) updateMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := getUUID(c)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		var user input.UserReq
		if err := c.BindJSON(&user); err != nil {
			ReturnErr(c, e.System.BadRequest)
			return
		}
		user.BindUUID(uuid)
		res, err := uh.userUsecase.UpdateUser(&user)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		ReturnSuccessData(c, &res)
	}
}

func (uh *userHandler) changeEmailNotify() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := getUUID(c)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		res, err := uh.userUsecase.ChangeEmailNotify(uuid)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		ReturnSuccessData(c, &res)
	}
}

func (uh *userHandler) deleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		uuid, err := entity.GetUUID(id)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		pass := c.Query("pass")
		if !entity.CheckPass(pass) {
			ReturnErr(c, e.System.BadRequest)
			return
		}
		err = uh.userUsecase.DeleteUser(uuid)
		if err != nil {
			ReturnErr(c, err)
			return
		}
		ReturnSuccess(c)
	}
}

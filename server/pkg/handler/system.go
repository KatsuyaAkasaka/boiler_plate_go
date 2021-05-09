package handler

import (
	"os"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/middleware"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type systemHandler struct {
	router        *gin.RouterGroup
	systemUsecase *usecase.SystemUsecase
}

func newSystemHandler(r *gin.RouterGroup, uc *usecase.Usecases) {
	sh := &systemHandler{
		systemUsecase: uc.System,
	}
	withPassRouterForGet := r.Group(
		"", func(c *gin.Context) {
			passStr := c.Query("pass")
			if !entity.CheckPass(passStr) {
				ReturnErr(c, e.System.BadRequest)
				c.Abort()
				return
			}
			c.Next()
		},
	)
	r.GET("/health-check", sh.healthcheck())
	withPassRouterForGet.GET("/commit", sh.commitSha())
	withPassRouterForGet.GET("/uuid-jwt/:id", sh.getUuidJwt())
	withPassRouterForGet.GET("/email-jwt/:email", sh.getEmailJwt())
	// imageHandler := uploader.GetUploadImageHandler(*conf, *repos)
	r.POST("/uploadImage", sh.uploadImage())
}

func (sh *systemHandler) healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		ReturnSuccess(c)
	}
}

func (sh *systemHandler) commitSha() gin.HandlerFunc {
	return func(c *gin.Context) {
		ReturnSuccessData(c, os.Getenv("IMAGE_TAG"))
	}
}

func (sh *systemHandler) getUuidJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidStr := c.Param("id")
		if uuidStr == "" {
			ReturnErr(c, e.System.BadRequest)
			return
		}
		uuid := entity.UUID(uuidStr)
		newToken, err := middleware.GenerateUUIDJWT(&uuid)
		if err != nil {
			ReturnErr(c, err)
		}
		ReturnSuccessData(c, newToken)
	}
}

func (sh *systemHandler) getEmailJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailStr := c.Param("email")
		if emailStr == "" {
			ReturnErr(c, e.System.BadRequest)
			return
		}
		email := entity.Email(emailStr)
		newToken, err := middleware.GenerateEmailJWT(&email)
		if err != nil {
			ReturnErr(c, err)
		}
		ReturnSuccessData(c, newToken)
	}
}

func (sh *systemHandler) uploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			log.Error(err)
			ReturnErr(c, e.System.BadRequest)
		}

		res, er := sh.systemUsecase.UploadImage(fileHeader)
		if er != nil {
			ReturnErr(c, er)
			return
		}
		ReturnSuccessData(c, &res)
	}
}

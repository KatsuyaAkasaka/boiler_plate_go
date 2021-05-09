package handler

import (
	"fmt"
	"net/http"
	"time"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/middleware"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

type handlers struct {
	user *userHandler
}

func ReturnErr(c *gin.Context, err e.Err) {
	c.JSON(err.HttpStatusCode, gin.H{"error": err.ErrCode})
}

func ReturnSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func ReturnSuccessData(c *gin.Context, d interface{}) {
	c.JSON(http.StatusOK, &d)
}

func getVersionPath(ver string) string {
	return fmt.Sprintf("/%s", string(ver))
}

func Start(uc *usecase.Usecases, middles *middleware.Middles, conf *config.GatewayInfo) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	// CORS 対応
	router.Use(
		cors.New(cors.Config{
			// 許可したいHTTPメソッドの一覧
			AllowMethods: []string{
				"POST",
				"GET",
				"OPTIONS",
				"PUT",
				"PATCH",
				"DELETE",
				"*",
			},
			// 許可したいHTTPリクエストヘッダの一覧
			AllowHeaders: []string{
				"*",
				"Authorization",
			},
			// 許可したいアクセス元の一覧
			AllowOrigins: []string{"*"},
			MaxAge:       24 * time.Hour,
		}),
	)

	versionBasedRouter := router.Group(getVersionPath(conf.Version))

	newUserHandler(versionBasedRouter, uc, middles)
	newSystemHandler(versionBasedRouter, uc)
	router.Run(fmt.Sprintf(":%s", conf.Port))
	log.Info("server started")
}

func getUUID(c *gin.Context) (*entity.UUID, e.Err) {
	val, exists := c.Get(middleware.GetJWTParam(middleware.UUID))
	if exists == false {
		return nil, e.System.TokenInvalid
	}

	return entity.GetUUID(val.(string))
}

func getEmail(c *gin.Context) (*entity.Email, e.Err) {
	val, exists := c.Get(middleware.GetJWTParam(middleware.Email))
	if exists == false {
		return nil, e.System.TokenInvalid
	}

	return entity.GetUserEmail(val.(string))
}

func getExp(c *gin.Context) (int64, e.Err) {
	val, exists := c.Get(middleware.GetJWTParam(middleware.Exp))
	if exists == false {
		return 0, e.System.TokenInvalid
	}

	return val.(int64), nil
}

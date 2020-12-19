package handler

import (
	"fmt"
	"net/http"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

type handlers struct {
	user *userHandler
}

func returnErr(c *gin.Context, err e.Err) {
	errStruct := e.GetError(err)
	c.JSON(errStruct.HttpStatusCode, gin.H{"error": errStruct.ErrCode, "description": errStruct.Description})
}

func returnSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func getVersionPath(ver string) string {
	return fmt.Sprintf("/%s", string(ver))
}

func Start(uc *usecase.Usecases, port string, ver string) {
	router = gin.Default()
	newUserHandler(router.Group("/users").Group(getVersionPath(ver)), uc)
	router.Run(fmt.Sprintf(":%s", port))
}

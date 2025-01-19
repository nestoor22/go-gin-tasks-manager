package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRouteHelloWork(ctx *gin.Context) {
	response := map[string]string{"hello": "world"}

	ctx.JSON(http.StatusOK, response)
	fmt.Println("TestRouteHelloWork")
}

func AddUserRoutes(router *gin.Engine) *gin.Engine {
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	baseRouter.GET("/test", TestRouteHelloWork)
	return router
}

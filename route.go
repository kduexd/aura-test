package main

import (
	"aura-test/controller/restful/item"
	"aura-test/controller/restful/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "aura-test/api/docs"
)

// Routing restful router
func Routing(router *gin.Engine) {
	if mode := gin.Mode(); mode == gin.DebugMode {
		router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	apiV1Group := router.Group("v1")
	{
		// user login
		apiV1Group.POST("login", user.Login)

		// get item list
		apiV1Group.GET("item", item.List)
	}

}

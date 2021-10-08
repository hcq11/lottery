package api

import (
	_ "lottery/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"lottery/configs"
	"net/http"
)
// @title Lottery API
// @version 1.0
// @description Lottery API.

// @contact.name Eden

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
type HTTPSrv struct {
}

func (hs *HTTPSrv) Run() {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.StaticFS(configs.ImgPrefix, http.Dir("./assets/img/lottery"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/v1")
	{
		v1.GET("/prizes", GetPrizes)
		v1.GET("/winners",GetWinners)
		v1.POST("/prize", AddPrize)
		v1.PUT("/prize",EditPrize)
		v1.DELETE("/prize/:id", DelPrize)
		v1.GET("/notify", NotifySocket)

		v1.POST("/lottery", Lottery)
	}
	router.Run(configs.HttpSrvPort)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		var isAccess = true
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

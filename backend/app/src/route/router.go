package route

import (
	"net/http"

	authsdk "atkit/auth_sdk"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//オーナーグループ作成
	ownerg := router.Group("/owner")
	ownerg.Use(authsdk.AuthMiddleware())
	{
		//グループを取得するエンドポイント
		ownerg.GET("/groups", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		//グループを作成するエンドポイント
		ownerg.POST("/group", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		//グループを削除するエンドポイント
		ownerg.DELETE("/group", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	return router
}
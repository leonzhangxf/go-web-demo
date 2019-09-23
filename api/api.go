package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"leonzhangxf-api/article"
	"leonzhangxf-api/auth"
	"leonzhangxf-api/config"
	_ "leonzhangxf-api/docs"
	"net/http"
)

var Engine *gin.Engine

// @title leonzhangxf BE API
// @version 1.0
// @description leonzhangxf's site api.
// @contact.name leonzhangxf
// @contact.url https://leonzhangxf.com
// @contact.email leonzhangxf@gmail.com
// @BasePath /
func init() {
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.Default()

	articleApi := article.Api{}
	authApi := auth.Api{}

	// open api
	open := Engine.Group("/open")
	{
		openV1 := open.Group("/v1")
		{
			articles := openV1.Group("/articles")
			{
				articles.GET("", articleApi.GetOpenArticles)
				articles.GET(":id", articleApi.GetOpenArticleById)
			}
		}
	}

	// internal api
	api := Engine.Group("/api")
	{
		apiV1 := api.Group("/v1")
		{
			authentication := apiV1.Group("/auth")
			{
				authentication.POST("/login", authApi.Login)
				authentication.POST("/refresh_token", authApi.RefreshToken)
			}

			articles := apiV1.Group("/articles")
			{
				// Auth token check
				articles.Use(authApi.TokenCheck)
				articles.GET("", articleApi.GetArticles)
			}
		}
	}

	// Enable swagger api docs: swag init
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Use the ping to indicate the server's status.
	Engine.GET("/ping", Ping)
}

func Ping(context *gin.Context) {
	err := config.Db.Ping()
	if nil != err {
		context.JSON(http.StatusServiceUnavailable,
			&PingResponse{Status: http.StatusServiceUnavailable, Msg: "DB ERR"})
		return
	}
	context.JSON(http.StatusOK, &PingResponse{Status: http.StatusOK, Msg: "OK"})
}

type PingResponse struct {
	Status uint32 `json:"status"`
	Msg    string `json:"msg"`
}

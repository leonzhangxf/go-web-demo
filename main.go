package main

import (
	"github.com/gin-gonic/gin"
	"leonzhangxf-api/api"
	_ "leonzhangxf-api/api"
	"leonzhangxf-api/db"
	_ "leonzhangxf-api/db"
	"log"
)

func main() {
	log.Println("leonzhangxf api start")

	api.Engine.GET("/ping", Ping)
	_ = api.Engine.Run()
}

func Ping(context *gin.Context) {
	err := db.Db.Ping()
	if nil != err {
		context.JSON(503, &PingResponse{Status: 503, Msg: "DB ERR"})
		return
	}
	context.JSON(200, &PingResponse{Status: 200, Msg: "OK"})
}

type PingResponse struct {
	Status uint32 `json:"status"`
	Msg    string `json:"msg"`
}

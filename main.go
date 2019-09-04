package main

import (
	"leonzhangxf-api/api"
	_ "leonzhangxf-api/api"
	_ "leonzhangxf-api/db"
	"leonzhangxf-api/util"
	_ "leonzhangxf-api/util"
)

func main() {
	addr := ":8080"
	util.Log.Infoln("leonzhangxf api start.", "The addr is", addr)
	err := api.Engine.Run(addr)
	if nil != err {
		util.Log.Fatalln("leonzhangxf api start failed.")
	}
}

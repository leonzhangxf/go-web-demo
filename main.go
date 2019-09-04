package main

import (
	"leonzhangxf-api/api"
	_ "leonzhangxf-api/api"
	_ "leonzhangxf-api/db"
	"log"
)

func main() {
	log.Println("leonzhangxf api start")
	_ = api.Engine.Run()
}

package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"leonzhangxf-api/util"
	"os"
)

var Db *sqlx.DB

var dbConfig DataSourceConfig

func init() {
	const dbFileName = "db.toml"
	dir, _ := os.Getwd()
	dbFileLocation := fmt.Sprintf("%v%vconfig%v%v", dir, string(os.PathSeparator),
		string(os.PathSeparator), dbFileName)
	util.Log.Infoln("Current db dir is ", dbFileLocation)
	if _, err := toml.DecodeFile(dbFileLocation, &dbConfig); nil != err {
		util.Log.Fatalln("Parse the db config fail.", err)
	}
	util.Log.Infoln("The db config parsed success.")

	dbConfigFormat := "%v:%v@tcp(%v:%v)/%v?%v"
	dbParams := "loc=Asia%2FShanghai&parseTime=true"
	dbConfigStr := fmt.Sprintf(dbConfigFormat, dbConfig.Username, dbConfig.Password,
		dbConfig.Address, dbConfig.Port, dbConfig.Db, dbParams)

	var err error
	Db, err = sqlx.Connect("mysql", dbConfigStr)
	if nil != err {
		util.Log.Fatalln("Connect to the db fail.", err)
	}
}

type DataSourceConfig struct {
	Address  string `toml:"address"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Db       string `toml:"db"`
}

package db

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var Db *sqlx.DB

var dbConfig DatasourceConfig

func Init() {
	const dbFileName = "db.toml"
	dir, _ := os.Getwd()
	dbFileLocation := fmt.Sprintf("%v%vconf%v%v", dir, string(os.PathSeparator),
		string(os.PathSeparator), dbFileName)
	log.Printf("Current db dir is %v\n", dbFileLocation)
	if _, err := toml.DecodeFile(dbFileLocation, &dbConfig); nil != err {
		log.Fatalln(err)
	}
	log.Printf("The parsed db config is %v\n", dbConfig)

	dbConfigFormat := "%v:%v@tcp(%v:%v)/%v?%v"
	dbParams := "loc=Asia%2FShanghai&parseTime=true"
	dbConfigStr := fmt.Sprintf(dbConfigFormat, dbConfig.Username, dbConfig.Password,
		dbConfig.Address, dbConfig.Port, dbConfig.Db, dbParams)

	var err error
	Db, err = sqlx.Connect("mysql", dbConfigStr)
	if nil != err {
		log.Fatalln(err)
	}
}

type DatasourceConfig struct {
	Address  string `toml:"address"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Db       string `toml:"db"`
}

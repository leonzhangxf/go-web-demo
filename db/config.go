package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

var dbConfig *DatasourceConfig

func Init() {
	dbConfig = &DatasourceConfig{
		Address:  "localhost",
		Port:     "3306",
		Username: "root",
		Password: "root",
		Db:       "article",
	}

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
	Address  string
	Port     string
	Username string
	Password string
	Db       string
}

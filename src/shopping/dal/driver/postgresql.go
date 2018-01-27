package driver

import (
	"database/sql"
	"fmt"
	"log"
)

var dbCon *sql.DB = nil

type dbConf struct {
	Host string
	Port int
	User string
	Password string
	DbName string
}

var conf = dbConf{
	Host: "localhost",
	DbName: "postgres",
	Port: 5432,
	User: "postgres",
	Password: "postgres",
}

func InitDb() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DbName)

	db_init, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db_connect := db_init
	//defer db.Close()
	//testing connection
	err = db_connect.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected! ")
	return db_connect
}

func GetDB() *sql.DB {
	if dbCon == nil {
		dbCon = InitDb()
	}
	return dbCon
}

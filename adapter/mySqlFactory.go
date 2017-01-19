package adapter

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import (
	"fmt"
	"github.com/jinzhu/configor"
)

func GetDB() *sql.DB {
	var DbConfig = struct {
		User string
		Password string
		DbName string
		Host string
		Port string
	}{}
	configor.Load(&DbConfig, "dbConfig.yml")

	dataSource := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true", DbConfig.User, DbConfig.Password,
		DbConfig.DbName)

	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
	return db
}
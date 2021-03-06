package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDB *sql.DB

func InitMysql(username, password, host string) (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/zhenaiwang?charset=utf8", username, password, host)
	MysqlDB, err = sql.Open("mysql", dsn)
	return MysqlDB, err
}

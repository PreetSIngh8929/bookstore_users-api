package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PreetSIngh8929/boookstore_utils-go/logger"
	// "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
)

// const (
// 	mysql_users_username = "mysql_users_username"
// 	mysql_users_password = "mysql_users_password"
// 	mysql_users_host     = "mysql_users_host"
// 	mysql_users_schema   = "mysql_users_schema"
// )

var (
	Client *sql.DB
	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host     = os.Getenv(mysql_users_host)
	// schema   = os.Getenv(mysql_users_schema)
)

func init() {
	// datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "abcdef", "localhost:3306", "users_db")
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	log.Println("database connection established")
}

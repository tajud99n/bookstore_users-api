package users

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tajud99n/bookstore_users-api/utils"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsername = "mysql_username"
	mysqlPassword = "mysql_password"
	mysqlHost     = "mysql_host"
	mysqlSchema   = "mysql_schema"
)

var (
	Client *sql.DB

	username = utils.GoDotEnvVariable(mysqlUsername)
	password = utils.GoDotEnvVariable(mysqlPassword)
	host     = utils.GoDotEnvVariable(mysqlHost)
	schema   = utils.GoDotEnvVariable(mysqlSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("databases successfully configured")
}

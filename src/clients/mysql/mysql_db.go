package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //   DRIVER_NAME = mysql  // we are not using this but we need this to make connections
)

var (
	Client *sql.DB
)

func init() {
	connection_string := fmt.Sprintf("%s:%s@tcp(%s)/%s",

		"root",
		"",
		"127.0.0.1",
		"oauth",
	)

	var err error

	Client, err = sql.Open("mysql", connection_string)

	if err != nil {
		panic(err)
	}

	if errs := Client.Ping(); errs != nil {
		panic(errs)
	}
	fmt.Println(" SQL database successfully configured ")

}

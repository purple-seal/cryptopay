package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init () {
	db, err := sql.Open("mysql", "webuser:zhdlsvorxhfl@#dnjf^dlf@tcp(18.182.133.180:3306)/bpay")
	if err != nil {
		fmt.Println("Cannot connect to db")
	}
	fmt.Println("init main")

	DB = db
}
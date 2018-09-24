package models

import (
	"fmt"
	"gopkg.in/gorp.v2"
)

type Invoice struct {
	name string
}


type Invoices struct {
	Uuid string `db:"uuid"`
	Price int32 `db:"price"`
}

var dao *gorp.DbMap

func (i Invoice) FindAll () []Invoices {
	if dao == nil {
		dao = initDb()
	}
	var rows []Invoices
	_, err := dao.Select(&rows,"select uuid, price from invoices")
	if err != nil {
		fmt.Println("oops")
		fmt.Println(err)
		return nil
	}

	return rows
}

func initDb () *gorp.DbMap {
	fmt.Println("init invoice")
	dbmap := &gorp.DbMap{Db: DB, Dialect: gorp.MySQLDialect{}}
	dbmap.AddTableWithName(Invoices{}, "invoices")

	return dbmap
}
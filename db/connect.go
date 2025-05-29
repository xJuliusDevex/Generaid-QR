package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Connect struct {
	Mydb string
	db   string
	Sql  *sql.DB
}

func (a *Connect) Start() error {
	var err error
	a.Sql, err = sql.Open(a.Mydb, a.db)
	if err != nil {
		return err
	}
	err = a.Sql.Ping()
	if err != nil {
		return err
	}
	return nil
}
func NewConnect(mydb, db string) Connect {
	return Connect{
		Mydb: mydb,
		db:   db,
	}

}
func (a Connect) Close() {
	a.Sql.Close()
}

package sqlconn

// sql connection to test for the datalayer

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
	db   *sql.DB
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
		if err != nil {
			log.Fatal(err)
		}
		db.Ping()
	})
	return db
}

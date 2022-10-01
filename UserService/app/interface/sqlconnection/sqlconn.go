package sqlconnection

import (
	"SepFirst/UserService/config"
	"database/sql"
	"log"
	"time"
)

var (
	DBConn *sql.DB
)

func Init(cfg config.SConfig) {
	var err error
	DBConn, err = OpenConnection(cfg)
	if err != nil {
		PrintInfo(cfg, err)
		Init(cfg)
	}

	DBConn.SetMaxOpenConns(100)
	DBConn.SetMaxIdleConns(10)
	DBConn.SetConnMaxLifetime(time.Minute * 3)

	if err = DBConn.Ping(); err != nil {
		PrintInfo(cfg, err)
		Init(cfg)
	}

	log.Println("Successfully connected to database")
}

// OpenConnection opens a connection to the database.
func OpenConnection(cfg config.SConfig) (*sql.DB, error) {

	db, err := sql.Open("mysql", cfg.GetConfig(config.USER_SERVICE_HOST).(string))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PrintInfo(config config.SConfig, err error) {
	log.Println("Exit with error: ", err)
	log.Println("Reconnecting to database with info: ", config)
	time.Sleep(time.Second * 3)
}

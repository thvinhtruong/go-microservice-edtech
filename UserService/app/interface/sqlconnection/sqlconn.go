package sqlconnection

import (
	"SepFirst/UserService/config"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
	driver = "mysql"
	source = "root:root@tcp(localhost:3306)/sefiUserService?parseTime=true"
)

func init() {
	var err error

	// configuration := config.GetInstance()
	// fmt.Println(configuration.GetConfig(config.MYSQL_USERNAME))
	// DBConn, err = sql.Open("mysql", fmt.Sprintf(
	// 	"%v:%v@tcp(%v:%v)/%v?parseTime=true",
	// 	configuration.GetConfig(config.MYSQL_USERNAME),
	// 	configuration.GetConfig(config.MYSQL_PASSWORD),
	// 	configuration.GetConfig(config.USER_SERVICE_HOST),
	// 	configuration.GetConfig(config.USER_SERVICE_PORT),
	// 	"sefiUserService",
	// ))

	DBConn, err = sql.Open(driver, source)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	DBConn.SetMaxOpenConns(100)
	DBConn.SetMaxIdleConns(10)
	DBConn.SetConnMaxLifetime(time.Minute * 3)

	if err = DBConn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	log.Println("Connected to database")
}

func GetDB() *sql.DB {
	return DBConn
}

func CloseDB() {
	DBConn.Close()
}

func PrintInfo(config config.SConfig, err error) {
	log.Println("Exit with error: ", err)
	log.Println("Reconnecting to database with info: ", config)
	time.Sleep(time.Second * 3)
}

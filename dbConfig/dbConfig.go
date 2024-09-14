package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConnect() {

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println(mysqlInfo)
	var err error
	DB, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// err = DB.Ping()
	// if err != nil {
	// 	log.Fatalf("failed to ping database: %v", err)
	// }

	fmt.Println("Connection opened to database")
}

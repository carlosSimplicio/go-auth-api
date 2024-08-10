package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error


func Connect() {
	config := mysql.Config{
		User: os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: "mysql",
		Addr: "127.0.0.1:3306",
		Net: "tcp",
	}

	db, err = sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")
}
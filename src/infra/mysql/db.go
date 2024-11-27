package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var pool *sql.DB
var err error

type MySqlClient struct{}

func (c *MySqlClient) Connect() {
	config := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: "auth",
		Addr:   os.Getenv("MYSQL_ADDRESS"),
		Net:    "tcp",
	}

	pool, err = sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := pool.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")
}

func (c *MySqlClient) Select(query string, params ...any) (*sql.Rows, error) {
	rows, err := pool.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

func (c *MySqlClient) Exec(query string, params ...any) (operationResult sql.Result, err error) {
	result, err := pool.Exec(query, params...)
	return result, err
}

func Close() {
	pool.Close()
}

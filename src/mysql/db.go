package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error


func Connect() {
	config := mysql.Config{
		User: os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: "auth",
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

func Select[T interface{}](query string, params ...any) ([]any, error) {
	var data []any
	rows, _ := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row T	
		t := reflect.TypeOf(&row).Elem()
		v := reflect.ValueOf(&row).Elem()
		numField := t.NumField()
		pointers := make([]any, numField)

		for i := 0; i < numField; i++ {
			structField := v.Field(i)
			pointers[i] =  structField.Addr().Interface()
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}
		
		data = append(data, row)
	}

	if err:= rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func Exec(query string, params ...any) (operationResult sql.Result, err error) {
	result, err := db.Exec(query, params...)
	return result, err
}
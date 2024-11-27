package utils

import (
	"database/sql"
	"reflect"
)

func GetRowsValues[T interface{}](rows *sql.Rows) ([]T, error) {
	var data []T
	for rows.Next() {
		var row T
		t := reflect.TypeOf(&row).Elem()
		v := reflect.ValueOf(&row).Elem()
		numField := t.NumField()
		pointers := make([]any, numField)

		for i := 0; i < numField; i++ {
			structField := v.Field(i)
			pointers[i] = structField.Addr().Interface()
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}

		data = append(data, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

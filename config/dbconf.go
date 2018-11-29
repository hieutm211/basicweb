
package config

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	Host = "localhost"
	User = "hieutm"
	Password = "hieutm211"
	DBname = "basicweb"
)

func InitDB() (*sql.DB, error) {
	connStmt := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", User, Password, Host, DBname)
	db, err := sql.Open("postgres", connStmt)
	if err != nil {
		return nil, err
	}
	return db, nil
}


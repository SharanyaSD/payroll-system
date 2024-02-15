package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "0602"
	dbname   = "my_pgdb"
)

func InitializeDB() (*sqlx.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlconn)
	//defer db.Close()
	if err != nil {
		return nil, err
	}

	//close connection if ping fails - ensuring db conn is valid
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

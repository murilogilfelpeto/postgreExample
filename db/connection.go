package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgres driver for sql.Open_
	"github.com/murilogilfelpeto/postgreExample/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	dbUrl := fmt.Sprintf(""+
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	connection, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()
	return connection, err
}

func CloseConnection(connection *sql.DB) {
	err := connection.Close()
	if err != nil {
		fmt.Println("Error closing connection: %v", err)
	}
}

package models

import (
	"github.com/murilogilfelpeto/postgreExample/db"
)

func Save(todo Todo) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer db.CloseConnection(connection)

	sql := "INSERT INTO todos (title, description, is_done) VALUES ($1, $2, $3) RETURNING id"

	err = connection.QueryRow(sql, todo.Title, todo.Description, todo.IsDone).Scan(&id)

	return id, err
}

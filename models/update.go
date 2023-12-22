package models

import "github.com/murilogilfelpeto/postgreExample/db"

func Update(id int64, todo Todo) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer db.CloseConnection(connection)

	sql := "UPDATE todos SET title = $1, description = $2, is_done = $3 WHERE id = $4"

	result, err := connection.Exec(sql, todo.Title, todo.Description, todo.IsDone, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

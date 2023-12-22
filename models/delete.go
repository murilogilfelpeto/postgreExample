package models

import "github.com/murilogilfelpeto/postgreExample/db"

func Delete(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer db.CloseConnection(connection)

	sql := "DELETE FROM todos WHERE id = $1"

	result, err := connection.Exec(sql, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

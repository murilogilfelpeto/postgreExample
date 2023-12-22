package models

import "github.com/murilogilfelpeto/postgreExample/db"

func GetById(id int64) (todo Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer db.CloseConnection(connection)

	row := connection.QueryRow("SELECT id, title, description, is_done FROM todos "+
		"WHERE id = $1", id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)
	return
}

func GetAll() (todos []Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer db.CloseConnection(connection)
	rows, err := connection.Query("SELECT id, title, description, is_done FROM todos")
	if err != nil {
		return
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}

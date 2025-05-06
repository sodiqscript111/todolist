package models

import (
	"todolist/db"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`
}

func (t *Todo) AddTodo() error {
	query := `INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Title, t.Description, t.Completed)
	if err != nil {
		return err
	}
	return nil
}

func GetTodoId(int) ([]Todo, error) {
	query := `SELECT * FROM todos`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func EditTodo(
	id int,
	todo Todo,
) error {
	query := `UPDATE todos SET title=?, description=?, completed=? WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description, todo.Completed, id)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTodos() ([]Todo, error) {
	rows, err := db.DB.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

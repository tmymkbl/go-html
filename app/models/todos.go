package models

import (
	"go-sample-todo/app/logs"
	"go-sample-todo/utils"
	"time"
)

type Todo struct {
	ID          int
	Title       string
	Content     string
	UserID      int
	CompletedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (u *User) CreateTodo(title string, content string) (err error) {
	cmd := `insert into todos (
		title, 
		content, 
		user_id, 
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd, title, content, u.ID, time.Now())
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}

func (u *User) GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, title, content, user_id from todos
	where user_id = ? and id = ? and deleted_at is null`
	todo = Todo{}
	err = Db.QueryRow(cmd, u.ID, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.UserID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}

	return todo, err
}

func (u *User) GetTodoWithCompleted(id int) (todo Todo, err error) {
	cmd := `select id, title, content, user_id, completed_at, created_at from todos
	where id = ? and user_id = ? and deleted_at is null and completed_at is not null`
	todo = Todo{}
	err = Db.QueryRow(cmd, id, u.ID).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.UserID,
		&todo.CompletedAt,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.DeletedAt)

	return todo, err
}

func (u *User) GetTodos() (todos []Todo, err error) {
	cmd := `select id, title, content, user_id, completed_at, created_at from todos where user_id = ? and deleted_at is null and completed_at is null order by created_at desc`
	rows, err := Db.Query(cmd)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.UserID,
			&todo.CompletedAt,
			&todo.CreatedAt)
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, title, content, user_id, completed_at, created_at from todos where user_id = ? and deleted_at is null and completed_at is null order by created_at desc`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.UserID,
			&todo.CompletedAt,
			&todo.CreatedAt)

		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
		}
		// contentHtml := template.HTML(todo.Content)
		// todo.Content = string(contentHtml)
		// todo.Content = string(template.HTML(todo.Content))
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUserWithCompleted() (todos []Todo, err error) {
	cmd := `select id, title, content, user_id, completed_at, created_at from todos where user_id = ? and deleted_at is null order by created_at desc`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.UserID,
			&todo.CompletedAt,
			&todo.CreatedAt)

		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set title = ?, content = ?, updated_at = ? where id = ?`
	_, err = Db.Exec(cmd, t.Title, t.Content, time.Now(), t.ID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}

func (t *Todo) CompleteTodo() (err error) {
	cmd := `update todos set completed_at = ? where id = ? and user_id = ?`
	_, err = Db.Exec(cmd, time.Now(), t.ID, t.UserID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `update todos set deleted_at = ? where id = ? and user_id = ?`
	_, err = Db.Exec(cmd, time.Now(), t.ID, t.UserID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}

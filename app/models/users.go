package models

import (
	"go-sample-todo/app/logs"
	"go-sample-todo/utils"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err
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

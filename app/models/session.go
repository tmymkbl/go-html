package models

import (
	"errors"
	"log"
	"net/http"
	"time"
)

type Session struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
func GetSession(_ http.ResponseWriter, request *http.Request) (sess Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid, 
		name, 
		email, 
		user_id, 
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Name, u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, name, email, user_id, created_at
	 from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Name,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, name, email, user_id, created_at
	 from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Name,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

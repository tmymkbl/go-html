package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go-sample-todo/config"

	"go-sample-todo/app/logs"
	"go-sample-todo/utils"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func InitDb() {
	file, line, funcName := utils.GetCurrentInfo()
	logs.Log.Debug("DB access init ", "funcName", funcName, "file", file, "line", line)

	if config.Config.SQLDriver == "sqlite3" {
		Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	} else if config.Config.SQLDriver == "mysql" {
		// Db, err := sql.Open("mysql", "ユーザー名:パスワード@tcp(ホスト:ポート)/データベース名?parseTime=true&loc=Asia%2FTokyo")
		// Db, err = sql.Open(config.Config.SQLDriver,
		// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia/Tokyo",
		var con string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.Config.DbUser,
			config.Config.DbPasswd,
			config.Config.DbHost,
			config.Config.DbPort,
			config.Config.DbName) + "?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Debug("DB access init ", "con", con, "func", funcName, "file", file, "line", line)
		Db, err = sql.Open(config.Config.SQLDriver, con)
	}
	if err != nil {
		// log.Error("DB ACCESS ERROR", "error", err)
		// log.Fatalln("DB ACEESS ERROR")
		// log.Fatalln(err)
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("DB connect ", "error", err, "tableNameUser", tableNameUser, "funcName", funcName, "file", file, "line", line)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		created_at TIMESTAMP)`, tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		// log.Println(err.Error())
		// log.Fatalf("%s TABLE CREATE ERROR\n", tableNameUser)
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("DB connect ", "error", err, "tableNameUser", tableNameUser, "funcName", funcName, "file", file, "line", line)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		content TEXT,
		user_id INTEGER UNSIGNED ,
		created_at TIMESTAMP)`, tableNameTodo)

	_, err = Db.Exec(cmdT)
	if err != nil {
		// log.Fatalf("%s TABLE CREATE ERROR\n", tableNameTodo)
		// log.Fatalln(err.Error())
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("DB connect ", "error", err, "tableNameTodo", tableNameTodo, "funcName", funcName, "file", file, "line", line)
	}

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		user_id INTEGER UNSIGNED,
		created_at TIMESTAMP)`, tableNameSession)

	_, err = Db.Exec(cmdS)
	if err != nil {
		// log.Fatalf("%s TABLE CREATE ERROR\n", tableNameSession)
		// log.Fatalln(err.Error())
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("DB connect ", "error", err, "tableNameSession", tableNameSession, "funcName", funcName, "file", file, "line", line)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

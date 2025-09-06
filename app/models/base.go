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

func ConnectDB() (err error) {
	file, line, funcName := utils.GetCurrentInfo()
	logs.Log.Debug("Connect DB ", "funcName", funcName, "file", file, "line", line)

	if config.Config.SQLDriver == "sqlite3" {
		Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	} else if config.Config.SQLDriver == "mysql" {
		var con string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.Config.DbUser,
			config.Config.DbPasswd,
			config.Config.DbHost,
			config.Config.DbPort,
			config.Config.DbName) + "?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Debug("Connect DB ", "con", con, "func", funcName, "file", file, "line", line)
		// Db, err := sql.Open("mysql", "ユーザー名:パスワード@tcp(ホスト:ポート)/データベース名?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo")
		Db, err = sql.Open(config.Config.SQLDriver, con)
	}
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}

	return err
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

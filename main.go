package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
	"os"

	"go-sample-todo/app/controllers"
	"go-sample-todo/app/logs"
	"go-sample-todo/app/models"
	"go-sample-todo/config"
	"go-sample-todo/utils"
)

func init() {
	if config.Config.LogOutput != "stdout" {
		// ファイルを開く
		var err error
		logs.Logfile, err = os.OpenFile(config.Config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logs.Logfile, err = os.Create(config.Config.LogFile)
			if err != nil {
				panic("Failed to create log file: " + err.Error())
			}
		}
		switch config.Config.LogOutput {
		case "file":
			logs.SetLogfile()
		case "multi":
			logs.SetLogMulti()
		}
		return
	}
	logs.SetLogStd() // デフォルトは標準出力
}

func main() {
	// defer logs.Logfile.Close()
	// logs.Log.Debug("start ", "file", "main.go", "line", 45, "funcName", "main")
	file, line, funcName := utils.GetCurrentInfo()
	logs.Log.Debug("start ", "file", file, "line", line, "funcName", funcName)

	err := models.ConnectDB()
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		os.Exit(-1)
	}

	mux := controllers.SetRoute()

	logs.Log.Info("start Server ", "port", config.Config.Port)

	csrfP := http.NewCrossOriginProtection()

	if config.Config.ServerMode == "cgi" {
		err := cgi.Serve(csrfP.Handler(mux))
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	} else {
		err := http.ListenAndServe(":"+config.Config.Port, csrfP.Handler(mux))
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	}
}

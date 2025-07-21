package main

import (
	"fmt"
	"net/http"

	"go-sample-todo/app/controllers"
	"go-sample-todo/app/models"
	"go-sample-todo/config"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()

	fmt.Println("Server start on port", config.Config.Port)
	err := http.ListenAndServe(":"+config.Config.Port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

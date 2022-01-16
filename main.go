package main

import (
	"fmt"

	"sample_todo_comp/app/controllers"
	"sample_todo_comp/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}

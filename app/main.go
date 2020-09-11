package main

import (
	"todo-tree/app/infrastructure/controller"
	"todo-tree/app/infrastructure/datastore"
	userRepository "todo-tree/app/infrastructure/datastore/user"

	gin "github.com/gin-gonic/gin"
)

func main() {
	initRouter()
}

func initRouter() {
	router := gin.Default()

	connectionString := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	commandHandler := datastore.NewCommandHandler(connectionString)
	queryHandler := datastore.NewQueryHandler(connectionString)

	userCommand := userRepository.NewDBUserCommand(commandHandler)
	userQuery := userRepository.NewDBUserQuery(queryHandler)
	userController := controller.NewUserController(userCommand, userQuery)
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Get(c) })
	router.Run()
}

package main

import (
	"log"

	"todo-tree/infrastructure/controller"
	"todo-tree/infrastructure/datastore"
	userRepository "todo-tree/infrastructure/datastore/user"

	gin "github.com/gin-gonic/gin"
	dig "go.uber.org/dig"
)

func main() {
	container := initDIContainer()
	router := initRouter(container)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}

func initDIContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() (datastore.DBContext, error) {
		return datastore.DBContext{
			ConnectionString: "root:password@tcp(172.21.0.2:3306)/todo_tree?charset=utf8mb4&parseTime=True&loc=Local",
		}, nil
	})
	container.Provide(datastore.NewCommandHandler)
	container.Provide(datastore.NewQueryHandler)
	container.Provide(userRepository.NewDBUserCommand)
	container.Provide(userRepository.NewDBUserQuery)
	container.Provide(controller.NewUserController)
	return container
}

func initRouter(container *dig.Container) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctrl *gin.Context) {
		ctrl.String(200, "Hello Gin!")
	})
	router.POST("/users", func(c *gin.Context) {
		container.Invoke(func(ctrl controller.UserController) {
			ctrl.Create(c)
		})
	})
	router.GET("/users/:id", func(c *gin.Context) {
		container.Invoke(func(ctrl controller.UserController) {
			ctrl.Get(c)
		})
	})
	return router
}

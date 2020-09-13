package main

import (
	"todo-tree/app/infrastructure/controller"
	"todo-tree/app/infrastructure/datastore"
	userRepository "todo-tree/app/infrastructure/datastore/user"

	gin "github.com/gin-gonic/gin"
	dig "go.uber.org/dig"
)

func main() {
	container := initDIContainer()
	router := initRouter(container)
	router.Run()
}

func initDIContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() (datastore.DBContext, error) {
		return datastore.DBContext{
			ConnectionString: "root:password@tcp(172.21.0.2:3306)/sample_docker_compose?charset=utf8mb4&parseTime=True&loc=Local",
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

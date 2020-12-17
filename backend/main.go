package main

import (
	"log"
	"todo-tree/infrastructure/controller"
	controllerInterface "todo-tree/interface/controller"
	"todo-tree/infrastructure/datastore"
	"todo-tree/infrastructure/datastore/task"
	cors "github.com/gin-contrib/cors"
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
			ConnectionString: "root:password@tcp(db:3306)/todo_tree?charset=utf8mb4&parseTime=True&loc=Local",
		}, nil
	})
	container.Provide(datastore.NewCommandHandler)
	container.Provide(datastore.NewQueryHandler)
	container.Provide(repository.NewDBTaskCommand)
	container.Provide(repository.NewDBTaskQuery)
	container.Provide(controller.NewTaskController)
	return container
}

func initRouter(container *dig.Container) *gin.Engine {
	router := gin.Default()
	
	//cors
	config := cors.DefaultConfig();
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/todo", func(c *gin.Context) {
		container.Invoke(func(ctrl controllerInterface.ITaskController) {
			ctrl.GetList(c)
		})
	})
	router.GET("/todo/:id", func(c *gin.Context) {
		container.Invoke(func(ctrl controllerInterface.ITaskController) {
			ctrl.Get(c)
		})
	})
	router.POST("/todo", func(c *gin.Context) {
		container.Invoke(func(ctrl controllerInterface.ITaskController) {
			ctrl.Create(c)
		})
	})
	router.PUT("/todo", func(c *gin.Context) {
		container.Invoke(func(ctrl controllerInterface.ITaskController) {
			ctrl.Update(c)
		})
	})
	router.DELETE("/todo/:id", func(c *gin.Context) {
		container.Invoke(func(ctrl controllerInterface.ITaskController) {
			ctrl.Delete(c)
		})
	})
	return router
}

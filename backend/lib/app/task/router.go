package task

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/todo-tree/app"
	infra "github.com/todo-tree/infrastructure/task"
	abstruct "github.com/todo-tree/interface/task"
	"go.uber.org/dig"
)

func getContainer() *dig.Container {
	container := app.GetDIContainer()
	container.Provide(infra.NewDBTaskQuery)
	container.Provide(infra.NewDBTaskCommand)
	container.Provide(infra.NewTaskController)
	return container
}

//InitGetListRouter router
func InitGetListRouter() *gin.Engine {
	var router = app.GetRouter()
	var container = getContainer()
	router.GET("/todo/list", func(c *gin.Context) {
		err := container.Invoke(func(ctrl abstruct.ITaskController) {
			ctrl.GetList(c)
		})
		if err != nil {
			log.Println(err)
		}else{
			log.Println("container invoke error not found.")
		}
	})
	return router
}

//InitGetRouter router
func InitGetRouter() *gin.Engine {
	var router = app.GetRouter()
	var container = getContainer()
	router.GET("/todo", func(c *gin.Context) {
		err := container.Invoke(func(ctrl abstruct.ITaskController) {
			ctrl.Get(c)
		})
		if err != nil {
			log.Println(err)
		}else{
			log.Println("container invoke error not found.")
		}
	})
	return router
}

//InitPostRouter router
func InitPostRouter() *gin.Engine {
	var router = app.GetRouter()
	var container = getContainer()
	router.POST("/todo", func(c *gin.Context) {
		err := container.Invoke(func(ctrl abstruct.ITaskController) {
			ctrl.Create(c)
		})
		if err != nil {
			log.Println(err)
		}else{
			log.Println("container invoke error not found.")
		}
	})
	return router
}

//InitPutRouter router
func InitPutRouter() *gin.Engine {
	var router = app.GetRouter()
	var container = getContainer()
	router.PUT("/todo", func(c *gin.Context) {
		err := container.Invoke(func(ctrl abstruct.ITaskController) {
			ctrl.Update(c)
		})
		if err != nil {
			log.Println(err)
		}else{
			log.Println("container invoke error not found.")
		}
	})
	return router
}

//InitDeleteRouter router
func InitDeleteRouter() *gin.Engine {
	var router = app.GetRouter()
	var container = getContainer()
	router.DELETE("/todo", func(c *gin.Context) {
		err := container.Invoke(func(ctrl abstruct.ITaskController) {
			ctrl.Delete(c)
		})
		if err != nil {
			log.Println(err)
		}else{
			log.Println("container invoke error not found.")
		}
	})
	return router
}

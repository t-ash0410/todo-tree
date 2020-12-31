package app

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/todo-tree/infrastructure/datastore"
	"go.uber.org/dig"
)

//GetDIContainer 初期設定済みのDIコンテナを提供
func GetDIContainer() *dig.Container {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	container := dig.New()
	container.Provide(func() (datastore.DBContext, error) {
		return datastore.DBContext{
			ConnectionString: connectionString,
		}, nil
	})
	container.Provide(datastore.NewCommandHandler)
	container.Provide(datastore.NewQueryHandler)
	return container
}

//GetRouter 初期設定済みのルータを提供
func GetRouter() *gin.Engine {
	router := gin.Default()
	allowOrigins := os.Getenv("ALLOW_ORIGINS")

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			allowOrigins,
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
		},
		MaxAge: 600,
	}))
	return router
}

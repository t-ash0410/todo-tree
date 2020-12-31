require (
	github.com/aws/aws-lambda-go v1.19.1
	github.com/awslabs/aws-lambda-go-api-proxy v0.9.0
	github.com/gin-gonic/gin v1.6.3
	github.com/todo-tree/app v0.0.0-00010101000000-000000000000
	github.com/todo-tree/infrastructure v0.0.0-00010101000000-000000000000
	github.com/urfave/cli v1.22.1 // indirect
)

module hello-world

go 1.13

replace github.com/todo-tree/app => ../../../lib/app

replace github.com/todo-tree/entity => ../../../lib/entity

replace github.com/todo-tree/infrastructure => ../../../lib/infrastructure

replace github.com/todo-tree/interface => ../../../lib/interface

replace github.com/todo-tree/usecase => ../../../lib/usecase

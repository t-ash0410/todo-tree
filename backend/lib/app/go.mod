module app

go 1.13

require (
	github.com/awslabs/aws-lambda-go-api-proxy v0.9.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/todo-tree/app v0.0.0-00010101000000-000000000000
	github.com/todo-tree/infrastructure v0.0.0-00010101000000-000000000000
	github.com/todo-tree/interface v0.0.0-00010101000000-000000000000
	go.uber.org/dig v1.10.0
)

replace github.com/todo-tree/app => ./

replace github.com/todo-tree/entity => ../entity

replace github.com/todo-tree/infrastructure => ../infrastructure

replace github.com/todo-tree/interface => ../interface

replace github.com/todo-tree/usecase => ../usecase

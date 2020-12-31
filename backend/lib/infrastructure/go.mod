module infrastructure

go 1.13

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/todo-tree/entity v0.0.0-00010101000000-000000000000
	github.com/todo-tree/infrastructure v0.0.0-00010101000000-000000000000
	github.com/todo-tree/interface v0.0.0-00010101000000-000000000000
	github.com/todo-tree/usecase v0.0.0-00010101000000-000000000000
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.9
)

replace github.com/todo-tree/infrastructure => ./

replace github.com/todo-tree/entity => ../entity

replace github.com/todo-tree/interface => ../interface

replace github.com/todo-tree/usecase => ../usecase

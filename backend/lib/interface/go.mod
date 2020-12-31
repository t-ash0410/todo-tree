module interface

go 1.13

replace github.com/todo-tree/interface => ./

replace github.com/todo-tree/entity => ../entity

require (
	github.com/todo-tree/entity v0.0.0-00010101000000-000000000000
	github.com/todo-tree/interface v0.0.0-00010101000000-000000000000
)

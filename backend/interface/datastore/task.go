package datastore

import (
	"todo-tree/entity/task"
)

//ITaskCommand TaskCommandのインターフェース
type ITaskCommand interface {
	Create(u entity.Task) (id int, err error)
	Update(u entity.Task) (err error)
	Delete(id int) (err error)
}

//ITaskQuery TaskQueryのインターフェース
type ITaskQuery interface {
	GetAll() (tasks []entity.Task, err error)
	FindByID(id int) (task entity.Task, err error)
}

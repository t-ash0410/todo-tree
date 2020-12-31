package task

import (
	"github.com/todo-tree/interface/controller"
)

//ITaskController TaskControllerのインターフェース
type ITaskController interface {
	Create(c controller.Context)
	Get(c controller.Context)
	GetList(c controller.Context)
	Update(c controller.Context)
	Delete(c controller.Context)
}

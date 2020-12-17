package controller

//ITaskController TaskControllerのインターフェース
type ITaskController interface {
	Create(c Context)
	Get(c Context)
	GetList(c Context)
	Update(c Context)
	Delete(c Context)
}

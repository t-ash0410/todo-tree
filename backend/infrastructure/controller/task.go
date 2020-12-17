package controller

import (
	"fmt"
	"strconv"
	"todo-tree/entity/task"
	"todo-tree/interface/controller"
	"todo-tree/interface/datastore"
	"todo-tree/usecase"
)

//TaskController Taskに関する操作の窓口
type TaskController struct {
	Interactor usecase.TaskInteractor
}

//NewTaskController UserControllerのコンストラクタ
func NewTaskController(command datastore.ITaskCommand, query datastore.ITaskQuery) controller.ITaskController {
	return TaskController {
		Interactor: usecase.TaskInteractor{
			Command: command,
			Query:   query,
		},
	}
}

//Create Taskの追加
func (controller TaskController) Create(c controller.Context) {
	task := entity.Task{}
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(400, err)
		fmt.Println(err)
		return
	}
	err := controller.Interactor.Add(task)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, task)
}

//Get Taskの取得
func (controller TaskController) Get(c controller.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := controller.Interactor.FindByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, task)
}

//GetList 複数のTaskを取得
func (controller TaskController) GetList(c controller.Context) {
	tasks, err := controller.Interactor.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, tasks);
}

//Update Taskの編集
func (controller TaskController) Update(c controller.Context){
	task := entity.Task{}
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(400, err)
		fmt.Println(err)
		return
	}
	taskWithAuthor, err := controller.Interactor.Update(task)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, taskWithAuthor)
}

//Delete Taskの削除
func (controller TaskController) Delete(c controller.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.Interactor.Delete(id);
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nil)
}

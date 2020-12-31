package task

import (
	"log"
	"strconv"

	"github.com/todo-tree/entity/task"
	"github.com/todo-tree/interface/controller"
	abstruct "github.com/todo-tree/interface/task"
	"github.com/todo-tree/usecase"
)

//Controller Taskに関する操作の窓口
type Controller struct {
	Interactor usecase.TaskInteractor
}

//NewTaskController UserControllerのコンストラクタ
func NewTaskController(command abstruct.ITaskCommand, query abstruct.ITaskQuery) abstruct.ITaskController {
	return Controller{
		Interactor: usecase.TaskInteractor{
			Command: command,
			Query:   query,
		},
	}
}

//Create Taskの追加
func (controller Controller) Create(c controller.Context) {
	task := task.Task{}
	if err := c.ShouldBind(&task); err != nil {
		log.Panicln(err)
		c.JSON(400, err)
		return
	}
	id, err := controller.Interactor.Add(task)
	if err != nil {
		log.Panicln(err)
		c.JSON(500, err)
		return
	}
	task.Id = id
	c.JSON(200, task)
}

//Get Taskの取得
func (controller Controller) Get(c controller.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Panicln(err)
		c.JSON(500, err)
		return
	}
	task, err := controller.Interactor.FindByID(id)
	if err != nil {
		log.Panicln(err)
		c.JSON(500, err)
		return
	}
	c.JSON(200, task)
}

//GetList 複数のTaskを取得
func (controller Controller) GetList(c controller.Context) {
	tasks, err := controller.Interactor.GetAll()
	if err != nil {
		log.Panicln(err)
		c.JSON(500, err)
		return
	}
	c.JSON(200, tasks)
}

//Update Taskの編集
func (controller Controller) Update(c controller.Context) {
	task := task.Task{}
	if err := c.ShouldBind(&task); err != nil {
		log.Panicln(err)
		c.JSON(400, err)
		return
	}
	taskWithAuthor, err := controller.Interactor.Update(task)
	if err != nil {
		log.Panicln(err)
		c.JSON(500, err)
		return
	}
	c.JSON(200, taskWithAuthor)
}

//Delete Taskの削除
func (controller Controller) Delete(c controller.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := controller.Interactor.Delete(id)
	if err != nil {
		log.Panicln(err)
		c.JSON(500, err)
		return
	}
	c.JSON(200, nil)
}

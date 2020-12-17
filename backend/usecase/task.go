package usecase

import (
	"todo-tree/entity/task"
	"todo-tree/interface/datastore"
)

//TaskInteractor ユーザに関するビジネスロジックを定義
type TaskInteractor struct {
	Query   datastore.ITaskQuery
	Command datastore.ITaskCommand
}

//GetAll 全てのタスクを返却
func (interactor *TaskInteractor) GetAll() (tasks []entity.Task, err error) {
	tasks, err = interactor.Query.GetAll()
	return
}

//FindByID IDで検索して返却
func (interactor *TaskInteractor) FindByID(id int) (task entity.Task, err error) {
	task, err = interactor.Query.FindByID(id)
	return
}

//Add タスクの追加
func (interactor *TaskInteractor) Add(task entity.Task) (err error) {
	_, err = interactor.Command.Create(task)
	return
}

//Update タスクの編集
func (interactor *TaskInteractor) Update(task entity.Task) (err error) {
	err = interactor.Command.Update(task)
	return
}

//Delete タスクの削除
func (interactor *TaskInteractor) Delete(id int) (err error) {
	err = interactor.Command.Delete(id)
	return
}

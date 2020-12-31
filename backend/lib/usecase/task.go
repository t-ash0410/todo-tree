package usecase

import (
	entity "github.com/todo-tree/entity/task"
	abstruct "github.com/todo-tree/interface/task"
)

//TaskInteractor ユーザに関するビジネスロジックを定義
type TaskInteractor struct {
	Query   abstruct.ITaskQuery
	Command abstruct.ITaskCommand
}

//GetAll 全てのタスクを返却
func (interactor *TaskInteractor) GetAll() (tasks []entity.Task, err error) {
	tasks, err = interactor.Query.GetAll()
	return
}

//FindByID IDで検索して返却
func (interactor *TaskInteractor) FindByID(id int) (task entity.WithAuthor, err error) {
	task, err = interactor.Query.FindByID(id)
	return
}

//Add タスクの追加
func (interactor *TaskInteractor) Add(task entity.Task) (int, error) {
	id, err := interactor.Command.Create(task)
	if err != nil {
		return id, err
	}
	return id, nil
}

//Update タスクの編集
func (interactor *TaskInteractor) Update(task entity.Task) (taskWithAuthor entity.WithAuthor, err error) {
	err = interactor.Command.Update(task)
	if err != nil {
		taskWithAuthor = entity.WithAuthor{}
		return
	}
	taskWithAuthor, err = interactor.FindByID(task.Id)
	return
}

//Delete タスクの削除
func (interactor *TaskInteractor) Delete(id int) (err error) {
	err = interactor.Command.Delete(id)
	return
}

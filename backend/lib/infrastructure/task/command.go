package task

import (
	entity "github.com/todo-tree/entity/task"
	infra "github.com/todo-tree/infrastructure/datastore"
	abstruct "github.com/todo-tree/interface/task"
)

//DBTaskCommand Taskの変更系クエリを司る
type DBTaskCommand struct {
	infra.SQLCommandHandler
}

//NewDBTaskCommand DBTaskCommandのコンストラクタ
func NewDBTaskCommand(commandHandler infra.SQLCommandHandler) abstruct.ITaskCommand {
	return DBTaskCommand{
		commandHandler,
	}
}

//Create タスクの作成
func (repo DBTaskCommand) Create(task entity.Task) (id int, err error) {
	result := repo.Conn.Create(&task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.Id, nil
}

//Update タスクの更新
func (repo DBTaskCommand) Update(task entity.Task) (err error) {
	result := repo.Conn.Updates(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete タスクの削除
func (repo DBTaskCommand) Delete(id int) (err error) {
	result := repo.Conn.Where("Id = ?", id).Delete(&entity.Task{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

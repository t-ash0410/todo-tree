package repository

import (
	"todo-tree/entity/task"
	abstruct "todo-tree/interface/datastore"
	infra "todo-tree/infrastructure/datastore"
)

//DBTaskCommand Taskの変更系クエリを司る
type DBTaskCommand struct {
	infra.SQLCommandHandler
}

//NewDBTaskCommand DBTaskCommandのコンストラクタ
func NewDBTaskCommand(commandHandler infra.SQLCommandHandler) abstruct.ITaskCommand {
	return DBTaskCommand {
		commandHandler,
	}
}

//Create タスクの作成
func (repo DBTaskCommand) Create(task entity.Task) (id int, err error) {
	result := repo.Conn.Create(&task)
	if result.Error != nil {
		panic(result.Error)
	}
	return task.ID, nil
}

//Update タスクの更新
func (repo DBTaskCommand) Update(task entity.Task) (err error) {
	repo.Conn.First(&task)
	repo.Conn.Save(&task)
	return nil
}

//Delete タスクの削除
func (repo DBTaskCommand) Delete(id int) (err error) {
	repo.Conn.Delete(&entity.Task{}, id)
	return nil
}

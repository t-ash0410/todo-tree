package task

import (
	entity "github.com/todo-tree/entity/task"
	infra "github.com/todo-tree/infrastructure/datastore"
	abstruct "github.com/todo-tree/interface/task"
)

//DBTaskQuery Taskの参照系クエリを司る
type DBTaskQuery struct {
	infra.SQLQueryHandler
}

//NewDBTaskQuery DBTaskQueryのコンストラクタ
func NewDBTaskQuery(queryHandler infra.SQLQueryHandler) abstruct.ITaskQuery {
	return DBTaskQuery{
		queryHandler,
	}
}

//GetAll 全てのタスクを返却
func (query DBTaskQuery) GetAll() (tasks []entity.Task, err error) {
	rows, err := query.Conn.Query("SELECT id, name, description FROM tasks")
	defer rows.Close()
	if err != nil {
		return
	}

	tasks = make([]entity.Task, 0)
	for rows.Next() {
		var id int
		var name string
		var description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return tasks, err
		}
		task := entity.Task{
			Id:          id,
			Name:        name,
			Description: description,
		}
		tasks = append(tasks, task)
	}
	return
}

//FindByID TaskをIDで検索して返却
func (query DBTaskQuery) FindByID(id int) (task entity.WithAuthor, err error) {
	var name string
	var description string

	if err := query.Conn.QueryRow("SELECT name, description FROM tasks WHERE id = ?", id).Scan(&name, &description); err != nil {
		return task, err
	}

	task.Id = id
	task.Name = name
	task.Description = description
	task.Author.Id = 1
	task.Author.Name = "ashikawa"
	return
}

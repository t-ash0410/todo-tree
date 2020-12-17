package repository

import (
	"todo-tree/entity/task"
	abstruct "todo-tree/interface/datastore"
	infra "todo-tree/infrastructure/datastore"
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

	for rows.Next() {
			var id int
			var name string
			var description string
			if err := rows.Scan(&id, &name, &description); err != nil {
					return tasks, err
			}
			task := entity.Task {
				ID: id,
				Name: name,
				Description: description,
			}
			tasks = append(tasks, task)
	}
  return
}

//FindByID TaskをIDで検索して返却
func (query DBTaskQuery) FindByID(id int) (task entity.Task, err error) {
	var name string
	var description string
	
	if err := query.Conn.QueryRow("SELECT name, description FROM tasks WHERE id = ?", id).Scan(&name, &description); err != nil {
		return task, err
	}

	task.ID = id
	task.Name = name
	task.Description = description
	return
}

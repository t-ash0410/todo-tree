package entity

import (
	user "todo-tree/entity/user"
)

//TaskWithAuthor Authorを持つタスク
type TaskWithAuthor struct {
	Task
	Author user.User
}

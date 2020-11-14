package controller

//Context HTTPの標準的なコンテキスト
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}

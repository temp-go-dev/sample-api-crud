package service

import "github.com/temp-go-dev/sample-api-crud/model"

// GetMessage メッセージを返却
func GetMessage() string {

	var todo = model.Todo{}
	todo.Id = "1"

	return "Hellow, Porld!"
}

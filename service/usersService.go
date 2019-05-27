package service

import (
	"github.com/google/uuid"
	"github.com/temp-go-dev/sample-api-crud/model"
)

// GetMessage メッセージを返却
func GetUsers() string {

	var todo = model.Todo{}
	todo.Id = "1"

	return "Hellow, Porld!"
}

// CreateUser user作成
func CreateUser(user *model.User) {

	// UUID払い出し
	uuid := uuid.New()
	uuidStr := uuid.String()

	// 設定
	user.Id = uuidStr
}

package service

import (
	"github.com/google/uuid"
	"github.com/temp-go-dev/sample-api-crud/model"
)

// ReadMessage メッセージを返却
func ReadUsers() string {

	var todo = model.Todo{}
	todo.Id = "1"

	return "Hellow, Porld!"
}

// CreateUser user作成
func CreateUser(user *model.User) {
	//userValue := *user
	// UUID払い出し
	uuid := uuid.New()
	uuidStr := uuid.String()
	// 設定
	user.Id = uuidStr
	// TODO : 永続化系処理
}

// ReadUser user作成
func ReadUser(userId string) model.User {
	// TODO : 永続化層からPK検索
	user := model.User{}
	user.Name = "hoge"
	user.FirstName = "fuga"
	user.LastName = "piyo"
	return user
}

func UpdateUser(userId string, user *model.User) {

}

func DeleteUser(userId string, user *model.User) {

}

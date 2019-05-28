package service

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/temp-go-dev/sample-api-crud/model"
)

// SearchUsers ユーザー検索
func SearchUsers() ([]model.User, error) {

	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/sampledb")
	if err != nil {
		fmt.Println(err.Error())
		// TODO : この実装はいいのだろうか? あとで再考したほうがよい
		return []model.User{}, err
	}
	defer db.Close()

	var (
		id         string
		firstName  string
		lastName   string
		email      string
		passsword  string
		phone      string
		userStatus int
		version    int
	)

	var users []model.User

	rows, err := db.Query("SELECT id, first_name, last_name, email, password, phone, userStatus, version FROM user ORDER BY id")
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &email, &passsword, &phone, &userStatus, &version)
		if err != nil {
			fmt.Println(err)
		}
		user := model.User{}
		user.Id = id
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.Password = passsword
		user.Phone = phone
		user.UserStatus = userStatus
		user.Version = version
		users = append(users, user)
	}

	return users, err
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

	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/sampledb")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var (
		id        string
		firstName string
		lastName  string
		email     string
	)

	rows, err := db.Query("SELECT id, first_name, last_name, email FROM user ORDER BY RAND() LIMIT 1")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &email)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("->" + id + "/" + firstName + "/" + lastName + "/")
	user := model.User{}
	user.Id = id
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	return user
}

func UpdateUser(userId string, user *model.User) {

}

func DeleteUser(userId string, user *model.User) {

}

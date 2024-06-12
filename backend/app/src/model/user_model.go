package model

import (
	"github.com/google/uuid"
)

type User struct {
	UserID string `gorm:"primaryKey"` //グループID
	Utype  string //ユーザータイプ

	Name     string //ユーザー名
	Password string //パスワード

	IconPath string  //アイコンパス

	GroupID string //グループID
}

func CreateUser(utype, name, password string) error {
	//ユーザーIDを生成
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	//ユーザーを作成
	result := dbconn.Create(&User{
		UserID:   uid.String(),
		Utype:    utype,
		Name:     name,
		Password: password,
		IconPath: "default.png",
	})

	return result.Error
}

//ユーザーを取得
func GetUser(userID string) (User, error) {
	var user User

	//ユーザーを一人取得
	result := dbconn.Where(&User{
		UserID: userID,
	}).First(&user)

	return user, result.Error
}


//ユーザーを削除
func DeleteUser(userID string) error {
	//ユーザーを削除
	result := dbconn.Where(&User{
		UserID: userID,
	}).Delete(&User{})

	return result.Error
}
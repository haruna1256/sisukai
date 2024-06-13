package model

import (
	"io"
	"os"
	"path"

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

	//アイコンパス
	iconpath := uid.String() + ".png"

	//ユーザーを作成
	result := dbconn.Create(&User{
		UserID:   uid.String(),
		Utype:    utype,
		Name:     name,
		Password: password,
		IconPath: iconpath,
	})

	//ファイルコピー
	err = copyfile(icon_default,path.Join(icon_dir,iconpath))

	//エラー処理
	if err != nil {
		return err
	}

	return result.Error
}

func copyfile(src string, dst string) error {
	//ファイルを開く
	src_file,err := os.Open(src)

	//エラー処理
	if err != nil {
		return err
	}

	//ファイルを閉じる
	defer src_file.Close()

	//ファイルを開く
	dst_file,err := os.Create(dst)

	//エラー処理
	if err != nil {
		return err
	}

	//ファイルを閉じる
	defer dst_file.Close()

	//コピー
	_,err = io.Copy(dst_file,src_file)
	if err != nil {
		return err
	}

	return nil
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

//アイコンパス取得
func GetIcon(userID string) (string, error) {
	//ユーザー取得
	user, err := GetUser(userID)

	return user.IconPath, err
}

//アイコン変更
func ChangeIcon(userID string,IconStream io.Reader) error {
	//ユーザー取得
	user, err := GetUser(userID)

	//エラー処理
	if err != nil {
		return err
	}

	//ファイルを開く
	dst_file,err := os.Create(path.Join(icon_dir,user.IconPath))

	//エラー処理
	if err != nil {
		return err
	}

	//ファイルを閉じる
	defer dst_file.Close()

	//コピー
	_,err = io.Copy(dst_file,IconStream)
	if err != nil {
		return err
	}

	return nil
}
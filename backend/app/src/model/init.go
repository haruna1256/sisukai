package model

import "gorm.io/gorm"

var (
	//DB接続
	dbconn *gorm.DB = nil

	//アイコンディレクトリ
	icon_dir string = "./assets/UserIcons/"


	//アイコンデフォルトパス
	icon_default string = "./assets/UserIcons/default.png"
)


func Init() {
	//DB接続
	Init_DB()
}
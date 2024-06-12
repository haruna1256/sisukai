package model

import "github.com/google/uuid"

type Group struct {
	GroupID   string `gorm:"primaryKey"` //グループID
	GroupName string //グループ名

	OwnerID   string //オーナーID
	Users     []User `gorm:"foreignkey:GroupID"`
}

//グループを作成
func CreateGroup(groupName, ownerID string) error {
	//グループIDを生成
	gid, err := uuid.NewRandom()
	if err != nil {
		return err
	}


}
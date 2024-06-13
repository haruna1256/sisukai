package model

import (
	"atkit/util"

	"github.com/google/uuid"
)

type Group struct {
	GroupID   string `gorm:"primaryKey"` //グループID
	GroupName string //グループ名

	OwnerID string        //オーナーID
	Members []GroupMember `gorm:"foreignKey:GroupID"` //メンバー
}


// グループを作成
func CreateGroup(groupName, ownerID string) (string,error) {
	//グループIDを生成
	gid, err := uuid.NewRandom()
	if err != nil {
		return "",err
	}

	//グループID
	groupID := gid.String()

	//グループを作成
	result := dbconn.Create(&Group{
		GroupID:   groupID,
		GroupName: groupName,
		OwnerID:   ownerID,
	})

	//エラー処理
	if result.Error != nil {
		return "",result.Error
	}

	//グループ取得
	get_group, err := Get_Group(groupID)

	//エラー処理
	if err != nil {
		return "",err
	}

	//メンバー追加
	err = get_group.Add_Member(GroupMember{
		UserID: ownerID,
		Role:   "Owner",
		MemberID: util.GenID(),
	})

	//エラー処理
	if err != nil {
		return "",err
	}

	return groupID,result.Error
}

// グループを取得
func Get_Group(groupID string) (Group, error) {
	var group Group

	//グループを取得
	result := dbconn.Where(&Group{
		GroupID: groupID,
	}).First(&group)

	return group, result.Error
}

// グループ名を変更
func (group *Group) Change_Name(groupName string) error {
	//グループ名を変更
	result := dbconn.Model(&group).Update("group_name", groupName)

	return result.Error
}

// メンバー一覧を取得
func (group *Group) Get_Members() ([]GroupMember, error) {
	var members []GroupMember

	//メンバーを取得
	err := dbconn.Model(&group).Association("Members").Find(&members)

	return members, err
}

// メンバーを取得
func (group *Group) Get_Member(userID string) (GroupMember, error) {
	var member GroupMember

	//メンバーを取得
	err := dbconn.Model(&group).Where(&GroupMember{UserID: userID}).Association("Members").Find(&member)

	return member, err
}

//メンバーを追加
func (group *Group) Add_Member(member GroupMember) error {
	//メンバーを追加
	err := dbconn.Model(&group).Association("Members").Append(&member)

	return err
}
package model

type GroupMember struct {
	MemberID string `gorm:"primaryKey"`
	GroupID string
	UserID  string
	Role    string
}

// ロール変更
func (member *GroupMember) Change_Role(role string) error {
	//ロールを変更
	result := dbconn.Model(&member).Update("role", role)

	return result.Error
}

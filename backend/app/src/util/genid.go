package util

import (
	"github.com/google/uuid"
)

//ID生成
func GenID() string {
	uid,err := uuid.NewRandom()

	if err != nil {
		return ""
	}

	return uid.String()
}
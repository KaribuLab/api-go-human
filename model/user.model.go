package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}

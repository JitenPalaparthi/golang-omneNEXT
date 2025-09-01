package models

import (
	"time"
)

type User struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordHash string `json:"-"` // omit from JSON
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified" gorm:"column:last_modified"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

func (u *User) Touch() {
	u.LastModified = time.Now().Unix()
}

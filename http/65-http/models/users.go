package models

import (
	"errors"
	"math/rand/v2"
	"time"
)

type User struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func NewUser(name, email string) *User {
	return &User{Id: uint(rand.IntN(100)), Name: name, Email: email, Status: "active", LastModified: time.Now().Unix()}
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("invalid email")
	}
	if u.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}

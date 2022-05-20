package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username       string
	HashedPassword string
	Role           string
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

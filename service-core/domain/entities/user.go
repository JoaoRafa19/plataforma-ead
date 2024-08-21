package entities

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	BaseEntity
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string 
}

func NewUser(Name string, Email string, Password string, Phone string) (*User, error) {
	user := &User{
		Name:     Name,
		Email:    Email,
		Password: Password,
		Phone:    Phone,
	}
	user.ID = uuid.New()

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("User Name invalid")
	}
	if u.Email == "" {
		return errors.New("User Email invalid")
	}
	if len(u.Password) < 8 {
		return errors.New("User Password invalid")
	}
	if u.Phone == "" {
		return errors.New("User Phone invalid")
	}
	return nil
}

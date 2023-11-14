package models

import (
	"errors"
	"mime/multipart"
	"regexp"
	"time"

	"github.com/google/uuid"
)

// role == 0 || 1 : user ? admin

// type_user == 0 || 1 : user ? teacher

// sex == 0 || 1 : men ? girl
type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"uniqueIndex;not null"`
	Email       string    `gorm:"not null"`
	Password    string    `gorm:"not null"`
	Role        int8      `gorm:"not null"`
	FullName    string    `gorm:"not null"`
	DateOfBirth string    `gorm:"not null"`
	Phone       string    `gorm:"not null"`
	Sex         int8      `gorm:"not null"`
	Address     string    `gorm:"not null"`
	City        string    `gorm:"not null"`
	District    string    `gorm:"not null"`
	Photo       string    `gorm:"not null"`
	TypeUser    int8      `gorm:"not null"`
	Deleted     int8      `gorm:"index"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

type SignUpInput struct {
	Name            string                `form:"name"`
	Email           string                `form:"email"`
	Password        string                `form:"password"`
	PasswordConfirm string                `form:"passwordConfirm"`
	Photo           *multipart.FileHeader `form:"photo"`
}

type SignInInput struct {
	Name     string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Otp      string `json:"otp"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      int8      `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Type      int8      `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignUpInputAdmin struct {
	Name            string                `form:"name"`
	Email           string                `form:"email"`
	Password        string                `form:"password"`
	PasswordConfirm string                `form:"passwordConfirm"`
	TypeUser        int8                  `form:"typeUser"`
	Role            int8                  `form:"role"`
	Photo           *multipart.FileHeader `form:"photo"`
}

type SignPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}

func ValidateUser(u SignUpInput) error {
	if u.Name == "" {
		return errors.New("name invalid!")
	}

	if u.Email == "" {
		return errors.New("email invalid!")
	}

	if u.Password == "" {
		return errors.New("password invalid!")
	}

	if len(u.Name) < 8 || len(u.Name) > 24 {
		return errors.New("len name invalid!")
	}

	if len(u.Password) < 8 {
		return errors.New("password min 8 character!")
	}

	if u.PasswordConfirm == "" {
		return errors.New("name invalid!")
	}

	if !isEmailValid(u.Email) {
		return errors.New("email invalid!")
	}

	return nil
}

func ValidateUserAdmin(u SignUpInputAdmin) error {
	if u.Name == "" {
		return errors.New("name invalid!")
	}

	if u.Email == "" {
		return errors.New("email invalid!")
	}

	if u.Password == "" {
		return errors.New("password invalid!")
	}

	if len(u.Name) < 8 || len(u.Name) > 24 {
		return errors.New("len name invalid!")
	}

	if len(u.Password) < 8 {
		return errors.New("password min 8 character!")
	}

	if u.PasswordConfirm == "" {
		return errors.New("name invalid!")
	}

	if !isEmailValid(u.Email) {
		return errors.New("email invalid!")
	}

	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

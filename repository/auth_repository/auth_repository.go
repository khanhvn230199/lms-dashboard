package authrepository

import (
	"kosei-jwt/models"

	"gorm.io/gorm"
)

type Authrepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) Authrepository {
	return Authrepository{
		DB: db,
	}
}

func (s *Authrepository) GetByID(id string) error {
	var user models.User
	result := s.DB.First(&user, "id = ? and deleted = 0", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Authrepository) CreateUser(user models.User) error {
	result := s.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Authrepository) GetUserByName(name string) (models.User, error) {
	var user models.User
	result := s.DB.First(&user, "name = ? and deleted = 0", name)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (s *Authrepository) GetUserByEmail(name string) (models.User, error) {
	var user models.User
	result := s.DB.First(&user, "email = ? and deleted = 0", name)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

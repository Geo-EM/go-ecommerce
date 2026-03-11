package repository

import (
	"e-commerce/internal/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUserByID(userId uint) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	UpdateUser(userId uint, user *domain.User) (domain.User, error)
	DeleteUser(userId uint) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

//

func (userRepo userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := userRepo.db.Create(&user).Error
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return domain.User{}, errors.New("Failed to create user")
	}

	return user, nil
}

func (userRepo userRepository) FindUserByID(userId uint) (domain.User, error) {
	return domain.User{}, nil
}

func (userRepo userRepository) FindUserByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (userRepo userRepository) UpdateUser(userId uint, user *domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (userRepo userRepository) DeleteUser(userId uint) (bool, error) {
	return false, nil
}

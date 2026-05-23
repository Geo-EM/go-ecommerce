package repository

import (
	"e-commerce/internal/domain"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	UpdateUser(userId uint, user domain.User) (domain.User, error)

	FindUserByID(userId uint) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	FindAllUsers() ([]domain.User, error)

	DeleteUser(userId uint) (bool, error)

	CreateBankAccount(bankAccount domain.BankAccount) (domain.BankAccount, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

//

func (userRepo userRepository) CreateUser(userInput domain.User) (domain.User, error) {
	var newUser domain.User = userInput

	err := userRepo.db.Create(&newUser).Error
	if err != nil {
		return domain.User{}, errors.New("Failed to create user")
	}

	return newUser, nil
}

func (userRepo userRepository) UpdateUser(userId uint, userInput domain.User) (domain.User, error) {
	var existingUser domain.User
	err := userRepo.db.Model(&existingUser).Clauses(clause.Returning{}).Where("id = ?", userId).Updates(userInput).Error
	if err != nil {
		return domain.User{}, errors.New("Failed to update user")
	}

	return existingUser, nil
}

func (userRepo userRepository) FindUserByID(userId uint) (domain.User, error) {
	var user domain.User
	err := userRepo.db.First(&user, userId).Error
	if err != nil {
		return domain.User{}, errors.New("User not found")
	}

	return user, nil
}

func (userRepo userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := userRepo.db.First(&user, "email = ?", email).Error
	if err != nil {
		return domain.User{}, errors.New("User not found")
	}

	return user, nil
}

func (userRepo userRepository) FindAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := userRepo.db.Find(&users).Error
	if err != nil {
		return []domain.User{}, errors.New("Failed to fetch users")
	}

	return users, nil
}

func (userRepo userRepository) DeleteUser(userId uint) (bool, error) {
	err := userRepo.db.Delete(&domain.User{}, userId).Error
	if err != nil {
		return false, errors.New("Failed to delete user")
	}

	return true, nil
}

func (userRepo *userRepository) CreateBankAccount(bankAccountInput domain.BankAccount) (domain.BankAccount, error) {
	var newBankAccount domain.BankAccount = bankAccountInput

	err := userRepo.db.Create(&newBankAccount).Error
	if err != nil {
		return domain.BankAccount{}, errors.New("Failed to create bank account")
	}

	return newBankAccount, nil
}

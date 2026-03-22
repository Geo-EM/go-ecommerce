package service

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/domain"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/repository"
	"errors"
)

type UserService struct {
	UserRepo     repository.UserRepository
	TokenService auth.TokenService
}

func (us UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := us.UserRepo.FindUserByEmail(email)
	return &user, err
}

func (us UserService) isUserVerified(userId uint) bool {
	user, err := us.UserRepo.FindUserByID(userId)
	return err == nil && user.Verified
}

func (us *UserService) RegisterUser(input userDto.RegisterUserDto) (string, error) {
	// Hash password
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return "", err
	}

	// Create user
	user, err := us.UserRepo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})
	if err != nil {
		return "", err
	}

	token, err := us.TokenService.GenerateToken(user.ID, user.Email, user.UserType)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) LoginUser(input userDto.LoginUserDto) (string, error) {
	user, err := us.findUserByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := auth.ValidatePassword(input.Password, user.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := us.TokenService.GenerateToken(user.ID, user.Email, user.UserType)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us UserService) GetVerificationCode(userId uint) (uint, error) {
	// If user already verified, ignore
	if us.isUserVerified(userId) {
		return 0, errors.New("user already verified")
	}

	// Generate verification code

	// Update user verification code

	return 0, nil
}

func (us UserService) VerifyUser(userId uint, code int) (bool, error) {
	return false, nil
}

func (us UserService) CreateUserProfile(userId uint, input any) (bool, error) {
	return false, nil
}

func (us UserService) GetUserProfile(userId uint) (*domain.User, error) {
	user, err := us.UserRepo.FindUserByID(userId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us UserService) UpdateUserProfile(userId uint, input any) (bool, error) {
	return false, nil
}

func (us UserService) UpdateUserCart(user domain.User, input any) (bool, error) {
	return false, nil
}

func (us UserService) GetUserCart(userId uint) ([]interface{}, error) {
	return nil, nil
}

func (us UserService) CreateOrder(user domain.User, input any) (bool, error) {
	return false, nil
}

func (us UserService) GetUserOrders(user domain.User) ([]interface{}, error) {
	return nil, nil
}

func (us UserService) GetUserOrderById(userId uint, orderId uint) (interface{}, error) {
	return nil, nil
}

func (us UserService) BecomeSeller(userId uint, input any) (bool, error) {
	return false, nil
}

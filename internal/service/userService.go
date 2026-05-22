package service

import (
	"e-commerce/config"
	"e-commerce/internal/auth"
	"e-commerce/internal/domain"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/repository"
	"e-commerce/pkg/notification"
	"errors"
	"fmt"
	"time"
)

type UserService struct {
	UserRepo     repository.UserRepository
	TokenService auth.TokenService
	AppConfig    config.AppConfig
}

// Private methods
func (us UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := us.UserRepo.FindUserByEmail(email)
	return &user, err
}

func (us UserService) isUserVerifiedById(userId uint) bool {
	user, err := us.UserRepo.FindUserByID(userId)
	return err == nil && user.Verified
}

func (us UserService) isUserVerified(user domain.User) bool {
	return user.Verified
}

func (us UserService) updateUserVerificationCode(userId uint, code uint) error {
	exp := time.Now().Add(30 * time.Minute)
	updatedUser := domain.User{VerificationCode: code, VerificationCodeExpiry: exp}
	_, err := us.UserRepo.UpdateUser(userId, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

//

// Public methods
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

func (us UserService) GetVerificationCode(userId uint) error {
	user, err := us.UserRepo.FindUserByID(userId)
	if err != nil {
		return err
	}

	// If user already verified, ignore
	if us.isUserVerified(user) {
		return errors.New("user already verified")
	}

	// Generate verification code
	code, err := auth.GenerateVerificationCode()
	if err != nil {
		return err
	}

	// Update user verification code
	if err := us.updateUserVerificationCode(userId, code); err != nil {
		return err
	}

	notificationClient := notification.NewNotificationClient(us.AppConfig)
	msg := fmt.Sprintf("(go-commerce): Your verification code is: %v", code)

	if err := notificationClient.SendSms(user.Phone, msg); err != nil {
		return err
	}

	return nil
}

func (us UserService) VerifyUserVerificationCode(userId uint, code uint) error {
	if us.isUserVerifiedById(userId) {
		return errors.New("user already verified")
	}

	user, err := us.UserRepo.FindUserByID(userId)
	if err != nil {
		return err
	}

	if !time.Now().Before(user.VerificationCodeExpiry) {
		return errors.New("verification code expired")
	}

	if user.VerificationCode == code {
		updatedUser := domain.User{Verified: true}
		_, err := us.UserRepo.UpdateUser(userId, updatedUser)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("invalid verification code")
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

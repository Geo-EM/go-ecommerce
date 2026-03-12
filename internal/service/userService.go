package service

import (
	"e-commerce/internal/domain"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/repository"
	"errors"
	"fmt"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (UserService) generateToken(user *domain.User) string {
	// TODO: Implement proper token generation logic here, possibly using JWT or another secure method
	return fmt.Sprintf("%v-%v-%v-token", user.ID, user.Email, user.UserType)
}

func (userService UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := userService.UserRepo.FindUserByEmail(email)
	return &user, err
}

func (userService UserService) RegisterUser(input userDto.RegisterUserDto) (string, error) {
	// TODO: Hash the password before storing it in the database
	user, err := userService.UserRepo.CreateUser(domain.User{
		Email: input.Email, Password: input.Password, Phone: input.Phone,
	})

	token := userService.generateToken(&user)
	return token, err
}

func (userService UserService) LoginUser(input userDto.LoginUserDto) (string, error) {
	user, err := userService.findUserByEmail(input.Email)

	// TODO: Hashed password comparison
	if err != nil || user.Password != input.Password {
		return "", errors.New("invalid credentials")
	}

	token := userService.generateToken(user)

	return token, err
}

func (userService UserService) GetVerificationCode(input any) (int, error) {
	return 0, nil
}

func (userService UserService) VerifyUser(userId uint, code int) (bool, error) {
	return false, nil
}

func (userService UserService) CreateUserProfile(userId uint, input any) (bool, error) {
	return false, nil
}

func (userService UserService) GetUserProfile(userId uint) (*domain.User, error) {
	return nil, nil
}

func (userService UserService) UpdateUserProfile(userId uint, input any) (bool, error) {
	return false, nil
}

func (userService UserService) UpdateUserCart(user domain.User, input any) (bool, error) {
	return false, nil
}

func (userService UserService) GetUserCart(userId uint) ([]interface{}, error) {
	return nil, nil
}

func (userService UserService) CreateOrder(user domain.User, input any) (bool, error) {
	return false, nil
}

func (userService UserService) GetUserOrders(user domain.User) ([]interface{}, error) {
	return nil, nil
}

func (userService UserService) GetUserOrderById(userId uint, orderId uint) (interface{}, error) {
	return nil, nil
}

func (userService UserService) BecomeSeller(userId uint, input any) (bool, error) {
	return false, nil
}

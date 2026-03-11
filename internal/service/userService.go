package service

import (
	"e-commerce/internal/domain"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/repository"
	"fmt"
	"log"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (userService UserService) findUserByEmail(email string) (*domain.User, error) {
	log.Println(email)
	return nil, nil
}

func (userService UserService) RegisterUser(input userDto.RegisterUserDto) (token string, err error) {
	log.Println(input)

	user, err := userService.UserRepo.CreateUser(domain.User{
		Email: input.Email, Password: input.Password, Phone: input.Phone,
	})

	// TODO: Implement proper token generation logic here, possibly using JWT or another secure method
	token = fmt.Sprintf("%v-%v-%v-token", user.ID, user.Email, user.UserType)

	return token, err
}

func (userService UserService) LoginUser(input userDto.LoginUserDto) (string, error) {
	log.Println(input)
	return "", nil
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

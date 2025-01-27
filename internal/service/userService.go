package service

import (
	"GoCart/internal/domain"
	"GoCart/internal/dto"
	"GoCart/internal/repository"
	"errors"
	"fmt"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform some db operations
	// write buisness logic
	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	log.Println(user)
	userInfo := fmt.Sprintf("%v,%v,%v", user.ID, user.Email, user.UserType)
	return userInfo, err
}

func (s UserService) Login(email, password string) (string, error) {
	user, err := s.Repo.FindUser(email)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("incorrect password")
	}
	return user.Email, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) (int, error) {

	return 0, nil
}

func (s UserService) CreateProfile(id uint, input interface{}) error {

	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input interface{}) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateCart(input interface{}, u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {

	return nil, nil
}

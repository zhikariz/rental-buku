package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	IsEmailAvailable(email string) (bool, error)
	Login(input LoginUserInput) (User, error)
	SavePhoto(id int, path string) (User, error)
	GetUserById(id int) (User, error)
	ResetPassword(input ResetPasswordInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Address = input.Address
	user.Email = input.Email
	user.Role = "Member"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) IsEmailAvailable(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SavePhoto(id int, path string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Photo = path

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found with that id")
	}

	return user, nil
}

func (s *service) ResetPassword(input ResetPasswordInput) (bool, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.MinCost)

	if err != nil {
		return false, err
	}

	user.Password = string(passwordHash)
	_, err = s.repository.Update(user)

	if err != nil {
		return false, err
	}

	return true, nil
}

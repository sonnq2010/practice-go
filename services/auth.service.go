package services

import (
	"practice-go/models"
	"practice-go/repos"
)

type IAuthService interface {
	SignUp() (models.User, error)
	SignIn() (models.User, error)
}

type AuthService struct {
	userRepo repos.IUserRepo
}

func NewAuthService(r repos.IUserRepo) IAuthService {
	return AuthService{
		userRepo: r,
	}
}

func (s AuthService) SignUp() (models.User, error) {
	user, err := s.userRepo.Create()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s AuthService) SignIn() (models.User, error) {
	user, err := s.userRepo.GetByID()
	if err != nil {
		return user, err
	}

	return user, nil
}

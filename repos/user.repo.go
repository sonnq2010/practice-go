package repos

import "practice-go/models"

type IUserRepo interface {
	GetAll() ([]models.User, error)
	GetByID() (models.User, error)
	Create() (models.User, error)
	Update() (models.User, error)
	Delete() error
}

type UserRepo struct{}

func NewUserRepo() IUserRepo {
	return UserRepo{}
}

func (r UserRepo) GetAll() ([]models.User, error) {
	var user models.User

	return []models.User{user}, nil
}

func (r UserRepo) GetByID() (models.User, error) {
	var user models.User

	return user, nil
}

func (r UserRepo) Create() (models.User, error) {
	var user models.User

	return user, nil
}

func (r UserRepo) Update() (models.User, error) {
	var user models.User

	return user, nil
}

func (r UserRepo) Delete() error {

	return nil
}

package api

import (
	"errors"
	"strings"
)

// UserService contains methods of the user service
type UserService interface {
	New(user NewUserRequest) error
}

// UserRepository is what lets our service do db operations without knowing anything about the implementations
type UserRepository interface {
	CreateUser(request NewUserRequest) error
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}

func (u *userService) New(user NewUserRequest) error {
	//validations
	if user.Email == "" {
		return errors.New("email required")
	}
	if user.Name == "" {
		return errors.New("name required")
	}

	if user.WeightGoal == "" {
		return errors.New("user service - weight goal required")
	}

	//normalisation
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	//create user from db
	err := u.storage.CreateUser(user)

	if err != nil {
		return err
	}
	return nil
}

package services

import (
	"github.com/tajud99n/bookstore_users-api/domain/users"
	"github.com/tajud99n/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser() (*users.User, *errors.RestErr) {
	return nil, nil
}

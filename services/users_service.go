package services

import (
	"fmt"
	"gin-rest-microservice/domain/httperrors"
	"gin-rest-microservice/domain/users"
)

var(
	UsersService = usersService{}

	registeredUsers = map[int64]*users.User{}
	currentUserId int64 = 1
)

type usersService struct {}

func (service usersService) Create(user users.User) (*users.User, *httperrors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Id = currentUserId
	currentUserId ++
	registeredUsers[user.Id] = &user

	return &user,nil
}

func (service usersService) Get(userId int64) (*users.User, *httperrors.HttpError) {
	if user := registeredUsers[userId]; user != nil {
		return user, nil
	}

	return nil, httperrors.NewNotFoundError(fmt.Sprintf("user with id %d not found", userId))
}
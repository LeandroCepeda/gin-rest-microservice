package users

import "gin-rest-microservice/domain/httperrors"

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

func (user User) Validate() *httperrors.HttpError {
	if user.FirstName == ""  {
		return httperrors.NewBadRequestError("invalid user first name")
	}
	if user.LastName == "" {
		return httperrors.NewBadRequestError("invalid user last name")
	}
	if user.Email == "" {
		return httperrors.NewBadRequestError("invalid email address")
	}
	return nil
}

package validator

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/saeidraei/go-jwt-auth/domain"
	"github.com/saeidraei/go-jwt-auth/uc"
)

type userValidator struct{}

func New() uc.UserValidator {
	return userValidator{}
}

func (userValidator) CheckUser(user domain.User) error {
	if ok := govalidator.IsEmail(user.Email); !ok {
		return errors.New("invalid email")
	}

	return nil
}

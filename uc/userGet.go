package uc

import (
	"github.com/saeidraei/go-jwt-auth/domain"
)

func (i interactor) UserGet(email string) (*domain.User, string, error) {
	user, err := i.userRW.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", ErrNotFound
	}
	if user.Email != email {
		return nil, "", errWrongUser
	}

	token, err := i.authHandler.GenUserToken(email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

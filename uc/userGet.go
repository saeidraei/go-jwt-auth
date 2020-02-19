package uc

import (
	"github.com/saeidraei/go-jwt-auth/domain"
)

func (i interactor) UserGet(userName string) (*domain.User, string, error) {
	user, err := i.userRW.GetByName(userName)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", ErrNotFound
	}
	if user.Name != userName {
		return nil, "", errWrongUser
	}

	token, err := i.authHandler.GenUserToken(userName)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

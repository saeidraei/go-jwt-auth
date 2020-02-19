package uc

import "github.com/saeidraei/go-jwt-auth/domain"

func (i interactor) UserLogin(email, password string) (*domain.User, string, error) {
	user, err := i.userRW.GetByEmailAndPassword(email, password)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", ErrNotFound
	}

	token, err := i.authHandler.GenUserToken(user.Name)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

package uc

import (
	"github.com/saeidraei/go-jwt-auth/domain"
)

func (i interactor) ProfileGet(requestingUserName, userName string) (*domain.User, bool, error) {
	user, err := i.userRW.GetByName(userName)
	if err != nil {
		return nil, false, err
	}
	if user == nil {
		return nil, false, errProfileNotFound
	}

	if requestingUserName == "" {
		return user, false, nil
	}

	_, err = i.userRW.GetByName(requestingUserName)
	if err != nil {
		return nil, false, err
	}
	if user == nil {
		return nil, false, errProfileNotFound
	}

	return user, false, nil
}

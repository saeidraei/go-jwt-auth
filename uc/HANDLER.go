package uc

import (
	"log"

	"github.com/saeidraei/go-jwt-auth/domain"
)

type Handler interface {
	ProfileLogic
	UserLogic
}

type ProfileLogic interface {
	ProfileGet(requestingUserName, userName string) (profile *domain.User, follows bool, err error)
}

type UserLogic interface {
	UserCreate(username, email, password string) (user *domain.User, token string, err error)
	UserLogin(email, password string) (user *domain.User, token string, err error)
	UserGet(userName string) (user *domain.User, token string, err error)
	UserEdit(userName string, fieldsToUpdate map[domain.UserUpdatableProperty]*string) (user *domain.User, token string, err error)
}

type TagsLogic interface {
	Tags() ([]string, error)
}

type HandlerConstructor struct {
	Logger        Logger
	UserRW        UserRW
	UserValidator UserValidator
	AuthHandler   AuthHandler
}

func (c HandlerConstructor) New() Handler {
	if c.Logger == nil {
		log.Fatal("missing Logger")
	}
	if c.UserRW == nil {
		log.Fatal("missing UserRW")
	}
	if c.UserValidator == nil {
		log.Fatal("missing UserValidator")
	}
	if c.AuthHandler == nil {
		log.Fatal("missing AuthHandler")
	}

	return interactor{
		logger:        c.Logger,
		userRW:        c.UserRW,
		userValidator: c.UserValidator,
		authHandler:   c.AuthHandler,
	}
}

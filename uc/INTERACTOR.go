package uc

import (
	"github.com/saeidraei/go-jwt-auth/domain"
)

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases and implement the Handler interface
type interactor struct {
	logger        Logger
	userRW        UserRW
	userValidator UserValidator
	authHandler   AuthHandler
}

// Logger : only used to log stuff
type Logger interface {
	Log(...interface{})
}

type AuthHandler interface {
	GenUserToken(email string) (token string, err error)
	GetUserName(token string) (email string, err error)
}

type UserRW interface {
	Create(username, email, password string) (*domain.User, error)
	GetByName(userName string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetByEmailAndPassword(email, password string) (*domain.User, error)
	Save(user domain.User) error
}

type UserValidator interface {
	CheckUser(user domain.User) error
}

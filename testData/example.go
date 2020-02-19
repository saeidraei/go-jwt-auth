package testData

import (
	"github.com/saeidraei/go-jwt-auth/domain"
)

var rickBio = "Rick biography string"
var janeImg = "jane img link"

func User(name string) domain.User {
	switch name {
	case "rick":
		return rick
	default:
		return jane
	}
}

var rick = domain.User{
	Name:      "rick",
	Email:     "rick@what.com",
	Bio:       &rickBio,
	ImageLink: nil,
	Password:  "password",
}

var jane = domain.User{
	Name:      "johnjacob",
	Email:     "joe@what.com",
	Bio:       nil,
	ImageLink: &janeImg,
	Password:  "password",
}

const TokenPrefix = "Token "

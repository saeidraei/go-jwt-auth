package uc

import "errors"

var (
	ErrAlreadyInUse = errors.New("this email is already in use")
	//ErrUserEmailAlreadyInUsed = errors.New("this email address is already in use")
	errWrongUser = errors.New("woops, wrong user")
	ErrNotFound  = errors.New("user not found")
)

package domain

import (
	"time"
)

// User represents a user account in the system
type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Bio       *string
	ImageLink *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type UserUpdatableProperty int

const (
	UserEmail UserUpdatableProperty = iota
	UserName
	UserBio
	UserImageLink
	UserPassword
)

func UpdateUser(initial *User, opts ...func(fields *User)) {
	for _, v := range opts {
		v(initial)
	}
}

func SetUserName(input *string) func(fields *User) {
	return func(initial *User) {
		if input != nil {
			initial.Name = *input
		}
	}
}

func SetUserEmail(input *string) func(fields *User) {
	return func(initial *User) {
		if input != nil {
			initial.Email = *input
		}
	}
}

// give empty string to delete it
func SetUserBio(input *string) func(fields *User) {
	return func(initial *User) {
		if input != nil {
			if *input == "" {
				initial.Bio = nil
				return
			}
			initial.Bio = input
		}
	}
}

// give empty string to delete it
func SetUserImageLink(input *string) func(fields *User) {
	return func(initial *User) {
		if input != nil {
			if *input == "" {
				initial.ImageLink = nil
				return
			}
			initial.ImageLink = input
		}
	}
}

func SetUserPassword(input *string) func(fields *User) {
	return func(initial *User) {
		if input != nil {
			initial.Password = *input
		}
	}
}

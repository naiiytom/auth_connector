package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound = errors.New("User not found")

	ErrInvalidID = errors.New("ID is not in its proper format")

	ErrAuthenticationFailure = errors.New("Authentication Failed")

	ErrForbidden = errors.New("Attempted action is not allowd")
)

func Create(n NewUser, now time.Time) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "generating password hash")
	}

	u := User{
		ID:           uuid.New().String(),
		Name:         n.Name,
		Email:        n.Email,
		PasswordHash: hash,
		DateCreated:  now.UTC(),
		DateUpdated:  now.UTC(),
	}

	return &u, nil
}

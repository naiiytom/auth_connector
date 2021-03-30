package user

import (
	"time"

	"gopkg.in/square/go-jose.v2/jwt"
)

type User struct {
	ID           string    `db:"user_id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash []byte    `db:"password_hash" json:"-"`
	DateCreated  time.Time `db:"date_created" json:"date_created"`
	DateUpdated  time.Time `db:"date_updated" json:"date_updated"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type WebUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	IsAdmin  string `json:"admin"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JwtUserToken struct {
	WebUser
	jwt.Claims
}

type JwtUserTokenResponse struct {
	Token string `json:"token"`
}

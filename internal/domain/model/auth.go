package model

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserID  int64 `json:"user_id"`
	IsAdmin bool  `json:"is_admin"`
	jwt.StandardClaims
}

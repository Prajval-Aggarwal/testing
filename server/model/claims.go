package model

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	Id   uuid.UUID
	Role string
	jwt.RegisteredClaims
}

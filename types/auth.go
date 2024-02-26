package types

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type InputRegister struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Birthday  string `json:"birthday" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Nik       string `json:"nik" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Role      string `json:"role" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JwtToken struct {
	Sub  uuid.UUID `json:"sub"`
	Role string    `json:"role"`
	jwt.RegisteredClaims
}

package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/draco121/common/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCustomClaims struct {
	Email     string             `json:"email"`
	UserId    primitive.ObjectID `json:"userId"`
	Role      constants.Role     `json:"role"`
	SessionId primitive.ObjectID `json:"sessionId"`
}

// DefaultClaims represents the default claims for the JWT token.
type DefaultClaims struct {
	JwtCustomClaims
	jwt.StandardClaims
}

// RefreshTokenClaims represents the claims for the refresh token.
type RefreshTokenClaims struct {
	SessionId primitive.ObjectID `json:"sessionId"`
	jwt.StandardClaims
}

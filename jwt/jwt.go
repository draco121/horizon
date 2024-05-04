package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/draco121/common/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"

	"github.com/draco121/common/models"
)

// JWTSecretKey is a secret key used to sign the JWT tokens.
var JWTSecretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT creates a new JWT token with default claims.
func GenerateJWT(customClaims *models.JwtCustomClaims) (string, error) {
	utils.Logger.Debug("generating JWT token")
	claims := models.DefaultClaims{
		JwtCustomClaims: *customClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecretKey)
	if err != nil {
		utils.Logger.Error("error generating JWT token", "error: ", err)
		return "", err
	}
	utils.Logger.Debug("generated JWT token")
	return signedToken, nil
}

// GenerateRefreshToken creates a new refresh token.
func GenerateRefreshToken(sessionId primitive.ObjectID) (string, error) {
	utils.Logger.Debug("generating refresh token")
	claims := models.RefreshTokenClaims{
		SessionId: sessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString(JWTSecretKey)
	if err != nil {
		utils.Logger.Error("error generating refresh token", "error: ", err)
		return "", err
	}
	utils.Logger.Debug("generated refresh token")
	return refreshToken, nil
}

func VerifyRefreshToken(refreshToken string) (*models.RefreshTokenClaims, error) {
	utils.Logger.Debug("verifying refresh token")
	// Validate the refresh token and get its claims
	token, err := jwt.ParseWithClaims(refreshToken, &models.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		utils.Logger.Error("error parsing refresh token", "error", err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.RefreshTokenClaims)
	if !ok || !token.Valid {
		return claims, fmt.Errorf("invalid refresh token")
	}
	utils.Logger.Debug("verified refresh token")
	return claims, nil
}

func VerifyJwtToken(jwtToken string) (*models.DefaultClaims, error) {

	utils.Logger.Debug("verifying jwt token")
	token, err := jwt.ParseWithClaims(jwtToken, &models.DefaultClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		utils.Logger.Error("error parsing jwt token", "error", err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.DefaultClaims)
	if !ok || !token.Valid {
		return claims, fmt.Errorf("invalid jwt token")
	}
	utils.Logger.Debug("verified jwt token")
	return claims, nil
}

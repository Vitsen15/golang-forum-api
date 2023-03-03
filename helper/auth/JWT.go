package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"go_forum/main/entity/auth"
	"time"
)

func GenerateJWT(credentials *auth.Credentials, expires time.Time) (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &auth.Claims{
		Username: credentials.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

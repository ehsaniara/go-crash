package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ehsaniara/go-crash/config"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(config.AppConfig.App.JwtSecret)
}

// GenerateToken generate tokens used for auth
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		//EncodeMD5(username),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-crash",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseJwtToken parsing token
func ParseJwtToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

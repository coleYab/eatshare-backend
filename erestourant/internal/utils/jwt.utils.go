package utils

import (
	"fmt"
	"time"

	"github.com/coleYab/erestourant/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateJwtToken(email, userid string) (string, error) {
	expiresIn := config.Cfg.JWTExpiration
	secret := config.Cfg.JWTSecret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     email,
		"id":        userid,
		"expiresAt": time.Now().Add(expiresIn).Format("02 03:04:05 MST 2006"),
	})

	return token.SignedString([]byte(secret))
}

// TODO: more tests if it can successfully validate the expiration time of the
// generated jwt token with less and lesser code if you can do that then yes
func ParseJwtToken(token string) (uuid.UUID, error) {
	secret := config.Cfg.JWTSecret
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("unable to cast the jwt token")
	}

	expiry := claims["expiresAt"].(string)
	expiryTime, err := time.Parse("02 03:04:05 MST 2006", expiry)
	if time.Now().Before(expiryTime) || err != nil {
		return uuid.UUID{}, fmt.Errorf("user token has expired %v", expiryTime.Format("02 03:04:05 MST 2006"))
	}

	id := claims["id"]
	return uuid.Parse(id.(string))
}

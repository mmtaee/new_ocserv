package crypto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"ocserv/pkg/config"
	"time"
)

func GenerateAccessToken(userID string, expire int64) (string, error) {
	cfg := config.Get()
	claims := jwt.MapClaims{
		"sub": userID,
		"jti": ulid.Make().String(),
		"exp": expire,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

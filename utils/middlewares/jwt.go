package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GenerateToken(id uint, role string) string {
	claim := make(jwt.MapClaims)
	claim["authorized"] = true
	claim["id"] = id
	claim["role"] = role
	claim["expired"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Error("error on token signed string: ", err.Error())
		return "cannot generate token"
	}

	return str
}

func ExtractToken(c echo.Context) (id uint, role string) {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return uint(claim["id"].(float64)), claim["role"].(string)
	}

	return 0, ""
}

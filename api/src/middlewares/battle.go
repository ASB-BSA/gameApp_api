package middlewares

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type ClaimsWithBattle struct {
	jwt.StandardClaims
}

func IsBattle(c *fiber.Ctx) error {
	cookie := c.Cookies("battle")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithBattle{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "対戦情報がありません",
		})
	}

	return c.Next()
}

func GenerateBattleToken(id uint) (string, error) {
	payload := ClaimsWithBattle{}

	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().AddDate(10, 0, 0).Unix()

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}

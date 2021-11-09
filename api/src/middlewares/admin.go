package middlewares

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type AdminClaimsWithScope struct {
	jwt.StandardClaims
}

func IsAdmin(c *fiber.Ctx) error {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic("Error getting .env data!")
	}
	SecretKey := os.Getenv("ADMIN_SECKEY")

	cookie := c.Cookies("admin-jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}

func GenerateAdminJWT(id uint) (string, error) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic("Error getting .env data!")
	}
	SecretKey := os.Getenv("ADMIN_SECKEY")

	payload := ClaimsWithScope{}

	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().AddDate(10, 0, 0).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}

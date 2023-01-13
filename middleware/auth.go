package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const secretKey = "secret"

type ClaimsWithScope struct {
	jwt.RegisteredClaims
	Scope string
}

func IsAuthenticated(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	claims := ClaimsWithScope{}
	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthenticated"})
	}

	payLoad := token.Claims.(*ClaimsWithScope)

	isCustomer := strings.Contains(ctx.Path(), "/api/loan/customer")
	isAdmin := strings.Contains(ctx.Path(), "/api/loan/admin")

	if (payLoad.Scope == "admin" && isCustomer) || (payLoad.Scope == "customer" && !isCustomer) {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized"})

	}

	if (payLoad.Scope == "customer" && isAdmin) || (payLoad.Scope == "admin" && !isAdmin) {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized"})

	}

	return ctx.Next()
}

func GetUserLogin(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("jwt")

	claims := ClaimsWithScope{}

	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	payLoad := token.Claims.(*ClaimsWithScope)

	return payLoad.Subject, nil
}

func GenerateJWT(adminEmail string, scope string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	payLoad := ClaimsWithScope{}
	payLoad.Subject = adminEmail
	payLoad.ExpiresAt = jwt.NewNumericDate(expireTime)
	payLoad.Scope = scope

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad).SignedString([]byte(secretKey))

}

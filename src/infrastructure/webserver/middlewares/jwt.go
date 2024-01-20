package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func AuthenticationMiddleware(g *echo.Group, secret string) {
	g.Use(setBearerRule)
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret),
	}))
	g.Use(validateJWTClient)
}

func setBearerRule(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if strings.HasPrefix(tokenString, "Bearer") == false {
			c.Request().Header.Set("Authorization", "Bearer "+tokenString)
		}
		return next(c)
	}
}

func validateJWTClient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_data", map[string]interface{}(claims))
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
}

type TokenClaim struct {
	VID       int    `json:"vid"`
	Token     string `json:"token"`
	Pl        string `json:"pl"`
	UserAgent string `json:"user_agent"`
	DeviceID  string `json:"device_id"`
	jwt.StandardClaims
}

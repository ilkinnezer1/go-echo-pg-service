package admin

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Authorization")
		// Removes the bearer for parsing token correctly
		if len(accessToken) > 7 && strings.ToLower(accessToken[0:7]) == "bearer " {
			accessToken = accessToken[7:]
		}
		if accessToken == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Token is missing",
			})
		}
		adminToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method and key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token signing method")
			}
			return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
		})

		if err != nil || adminToken == nil {
			return echo.ErrUnauthorized
		}
		// Set the admin token in the context
		c.Set("admin", adminToken)

		admin, ok := c.Get("admin").(*jwt.Token)
		if !ok || admin == nil {
			return echo.ErrUnauthorized
		}

		claims, ok := admin.Claims.(jwt.MapClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		isAdmin, ok := claims["admin"].(bool)
		if !ok || !isAdmin {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

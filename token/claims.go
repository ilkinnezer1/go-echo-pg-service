package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
)

var _ = godotenv.Load()

type JwtAdminClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

var JwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(JwtAdminClaims)
	},
	SigningKey: []byte(os.Getenv("JWT_SIGNING_KEY")),
}

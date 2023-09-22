package handlers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main/config"
	"main/migrate"
	"main/models"
	"main/token"
	"net/http"
	"os"
	"time"
)

var db, _ = migrate.ConnectDatabase()

// Error messages
var errInvalidUserName = InvalidUserCredentials()
var errFailHashPassword = FailHashPasswd()
var errFailStoreData = FailStoreData()
var errFailUpdateLastLoginTime = FailUpdateLastLoginTime() // Make Connection with the database

func LoginHandler(c echo.Context) error {
	err := godotenv.Load()
	username, password := config.GetCredentials()

	// Get the username and password from the request body
	reqUsername := c.FormValue("username")
	reqPassword := c.FormValue("password")

	// Check if the username and password are correct
	if reqUsername != username || reqPassword != password {
		return c.JSON(http.StatusUnauthorized, errInvalidUserName)
	}
	// Hash the password using bcrypt before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errFailHashPassword)
	}

	user := &models.User{
		Username: reqUsername,
		Password: string(hashedPassword),
	}
	if err := db.Where("username = ?", reqUsername).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// User is logging in for the first time
			user.LastLoginTime = time.Now()
			if err := db.Create(user).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, errFailStoreData)
			}

		} else {
			return c.JSON(http.StatusInternalServerError, errFailStoreData)
		}
	} else {
		user.PrevLoginTime = user.LastLoginTime
		// Update Last Login Time
		user.LastLoginTime = time.Now()
		if err := db.Save(user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, errFailUpdateLastLoginTime)
		}
	}

	claims := &token.JwtAdminClaims{
		Name:  "Admin",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := generatedToken.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	ErrorHandler(err)

	response := map[string]interface{}{
		"token":           accessToken,
		"last_login_time": user.PrevLoginTime.Format(time.RFC3339),
	}

	// Return the last login time and access token
	return c.JSON(http.StatusOK, response)
}

package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("maycakepheh3")
var tokenName = "token"

type Claims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"usertype"`
	jwt.StandardClaims
}

func generateToken(c echo.Context, id int, name string, userType int) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	c.SetCookie(&http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})
}

func ResetUserToken(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
		Path:     "",
	})
}

func Authenticate(next echo.HandlerFunc, accessType int) echo.HandlerFunc {
	return func(c echo.Context) error {
		isValidToken := validateUserToken(c, accessType)
		if !isValidToken {
			return c.JSON(http.StatusUnauthorized, "Unauthorized User")
		} else {
			return next(c)
		}
	}
}

func validateUserToken(c echo.Context, accessType int) bool {
	isAccessTokenValid, id, name, userType := validateTokenFromCookies(c)
	log.Print(id, name, userType, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(c echo.Context) (bool, int, string, int) {
	if cookie, err := c.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		}
	}
	return false, -1, "", -1
}

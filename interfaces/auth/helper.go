package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

func Claims(c echo.Context) *domain.AuthClaims {
	return c.Get("user").(*jwt.Token).Claims.(*domain.AuthClaims)
}
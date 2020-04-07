package handler

import (
	"GolangEchoGorm/helper"
	"github.com/labstack/echo"
	"net/http"
)

func Restricted(c echo.Context) error {
	user := c.Get("user")
	identifier, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}

	if !valid {
		return c.String(http.StatusUnauthorized, "U are Logged out")
	}
	return c.String(http.StatusOK, "Welcome "+identifier.Username+"!")
}

func Logout(c echo.Context) error {

	status, err := helper.RemoveToken(c.Get("user"))
	if err != nil {
		return c.JSON(status, err)
	}

	return c.String(http.StatusOK, "Loget out!")
}

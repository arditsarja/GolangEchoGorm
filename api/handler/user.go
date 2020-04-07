package handler

import (
	"GolangEchoGorm/api/mapper"
	"GolangEchoGorm/helper"
	"GolangEchoGorm/model"
	"github.com/labstack/echo"
	"net/http"
)

// Handler
func CreateUser(c echo.Context) (err error) {
	user := c.Get("user")
	identifier, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, "U are Logged out")
	}
	if !identifier.Admin {
		return c.JSON(http.StatusUnauthorized, "U are Not Admin to creat user")
	}
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, u)
	}
	u.HashPassword()
	status, err := mapper.UserMapper{}.CreateUser(*u)

	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(http.StatusOK, u)
}

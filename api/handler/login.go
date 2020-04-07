package handler

import (
	"GolangEchoGorm/helper"
	"GolangEchoGorm/util"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type loginStruct struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

func Login(c echo.Context) error {
	login := new(loginStruct)

	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, existUser, err := helper.GetExistUser(login.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if !existUser {
		return c.JSON(http.StatusBadRequest, "User does not exist")
	}

	md5Password := util.CreateHash(login.Password)
	if login.Username != user.Username || md5Password != user.Password {
		return echo.ErrUnauthorized
	}
	fmt.Println(user)

	tokenResponse, err := helper.CreateJwtToken(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, tokenResponse)
	}
	return c.JSON(http.StatusOK, tokenResponse)
}

package handler

import (
	"GolangEchoGorm/api/mapper"
	"GolangEchoGorm/helper"
	"GolangEchoGorm/model"
	"github.com/labstack/echo"
	"net/http"
)

func CreateTable3(c echo.Context) error {
	user := c.Get("user")
	_, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, "U are Logged out")
	}
	table3 := new(model.Table3)
	if err = c.Bind(table3); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	table3Mapper := mapper.Table3Mapper{Table3: *table3}
	status, err := table3Mapper.CreateTable3()
	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(http.StatusOK, table3)
}

func GetTable3(c echo.Context) error {
	user := c.Get("user")
	nameOrId := c.Param("name_or_id")
	_, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, "U are Logged out")
	}
	status, err, response := mapper.Table3Mapper{}.GetTable3(nameOrId)
	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(http.StatusOK, response)
}


package handler

import (
	"GolangEchoGorm/api/mapper"
	"GolangEchoGorm/helper"
	"GolangEchoGorm/model"
	"github.com/labstack/echo"
	"net/http"
)

func CreateTable1(c echo.Context) error {
	user := c.Get("user")
	_, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, "U are Logged out")
	}
	table1 := new(model.Table1)
	if err = c.Bind(table1); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	table1Mapper := mapper.Table1Mapper{Table1Mapper: *table1}
	status, err := table1Mapper.CreateTable1()
	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(http.StatusOK, table1)
}
func GetTable1(c echo.Context) error {
	user := c.Get("user")
	name_or_id := c.Param("name_or_id")
	_, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, "U are Logged out")
	}
	status, err, response := mapper.Table1Mapper{}.GetTable1(name_or_id)
	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(http.StatusOK, response)
}

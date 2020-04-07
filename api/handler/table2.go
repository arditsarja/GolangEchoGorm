package handler

import (
	"GolangEchoGorm/api/mapper"
	"GolangEchoGorm/helper"
	"GolangEchoGorm/model"
	"github.com/labstack/echo"
	"net/http"
)

func CreateTable2(c echo.Context) error {
	user := c.Get("user")
	_, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, "U are Logged out")
	}
	table2 := new(model.Table2)
	if err = c.Bind(table2); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	table2Mapper := mapper.Table2Mapper{Table2: *table2}
	status, err := table2Mapper.CreateTable2()
	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(http.StatusOK, table2)
}

func GetTable2(c echo.Context) error {
	user := c.Get("user")
	nameOrId := c.Param("name_or_id")
	_, valid, err := helper.GetIdentifierFromJWT(user)
	if err != nil {
		return err
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, "U are Logged out")
	}
	status, err, response := mapper.Table2Mapper{}.GetTable2(nameOrId)
	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(http.StatusOK, response)
}


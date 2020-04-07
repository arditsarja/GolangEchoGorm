package main

import (
	"GolangEchoGorm/api/handler"
	"GolangEchoGorm/db"
	"GolangEchoGorm/helper"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := db.InitMigrations()
	if err != nil {
		fmt.Println(err)
		return
	}
	config := middleware.JWTConfig{
		Claims:        &helper.JwtCustomClaims{},
		SigningMethod: "HS512",
		SigningKey:    []byte("secret"),
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	e.POST("/login", handler.Login)

	r := e.Group("")

	r.Use(middleware.JWTWithConfig(config))

	r.GET("", handler.Restricted)
	r.GET("/logout", handler.Logout)
	r.GET("/create_user", handler.CreateUser)
	r.POST("/create_user", handler.CreateUser)

	r.POST("/table1", handler.CreateTable1)
	r.GET("/table1/:name_or_id", handler.GetTable1)

	r.POST("/table2", handler.CreateTable2)
	r.GET("/table2/:name_or_id", handler.GetTable2)

	r.POST("/table3", handler.CreateTable3)
	r.GET("/table3/:name_or_id", handler.GetTable3)

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}

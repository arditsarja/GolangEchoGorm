package db

import (
	"GolangEchoGorm/configuration"
	"GolangEchoGorm/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //
	"github.com/jinzhu/gorm"
)

// Database function
// Returns gorm connection database
func Database() (*gorm.DB, error) {
	uri := configuration.DbConfig.DbUser + ":" + configuration.DbConfig.DbPassword +
		"@tcp(" + configuration.DbConfig.DbHost + ":" + configuration.DbConfig.DbPort + ")/" + configuration.DbConfig.DbName
	db, err := gorm.Open("mysql", uri+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitMigrations function
// Initializes migrations in the database
func InitMigrations() error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	dbResponse := db.AutoMigrate(

		&model.Table3{},
		&model.Table1{},
		&model.Table2{},
		&model.Token{},
		&model.User{},
	)

	if dbResponse.Error != nil {
		fmt.Println(dbResponse.Error)
		return dbResponse.Error
	}

	username := "admin"
	adminUserExist := model.User{Username: username}
	db.First(&adminUserExist)
	if adminUserExist.ID > 0 {
		//admin exist
		// dont create admin
		return nil
	}

	adminUser := model.User{
		Username:  configuration.AdminConfig.Username,
		FirstName: configuration.AdminConfig.FirstName,
		LastName:  configuration.AdminConfig.LastName,
		Password:  configuration.AdminConfig.Password,
		Email:     configuration.AdminConfig.Email,
		Address:   configuration.AdminConfig.Address,
		Admin:     configuration.AdminConfig.Admin,
	}
	adminUser.HashPassword()
	db.Save(&adminUser)
	return nil
}

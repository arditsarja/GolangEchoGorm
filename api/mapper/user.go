package mapper

import (
	"GolangEchoGorm/db"
	"GolangEchoGorm/model"
	"net/http"
)

type UserMapper struct {
}

func (UserMapper) CreateUser(user model.User) (int, error) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer currentDb.Close()

	dbResponse := currentDb.Save(&user)
	if dbResponse.Error != nil {
		return http.StatusInternalServerError, dbResponse.Error
	}

	return http.StatusOK, nil
}

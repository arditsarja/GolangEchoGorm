package helper

import (
	"GolangEchoGorm/db"
	"GolangEchoGorm/model"
)

func GetExistUser(username string) (model.User, bool, error) {
	user := model.User{Username: username}
	currentDb, err := db.Database()
	if err != nil {
		return user, false, err
	}

	currentDb.Where(&user).First(&user)
	if user.ID < 1 {
		return user, false, nil
	}

	return user, true, nil
}

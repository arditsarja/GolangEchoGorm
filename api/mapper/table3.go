package mapper

import (
	"GolangEchoGorm/db"
	"GolangEchoGorm/model"
	"net/http"
)

type Table3Mapper struct {
	Table3 model.Table3
}

func (t Table3Mapper) CreateTable3() (int, error) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer currentDb.Close()
	responseDB := currentDb.Save(&t.Table3)
	if responseDB.Error != nil {
		return http.StatusInternalServerError, responseDB.Error
	}
	return http.StatusOK, nil
}

func (Table3Mapper) GetTable3(nameOrId string) (int, error, interface{}) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err, nil
	}
	defer currentDb.Close()

	table3 := model.Table3{}

	responseDB := currentDb.Where(&table3).Find(&table3)
	if responseDB.Error != nil {
		return http.StatusInternalServerError, responseDB.Error, nil
	}

	return http.StatusOK, nil, table3
}

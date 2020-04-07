package mapper

import (
	"GolangEchoGorm/db"
	"GolangEchoGorm/model"
	"net/http"
)

type Table1Mapper struct {
	Table1Mapper model.Table1
}

func (m Table1Mapper) CreateTable1() (int, error) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer currentDb.Close()
	responseDB := currentDb.Save(&m.Table1Mapper)
	if responseDB.Error != nil {
		return http.StatusInternalServerError, responseDB.Error
	}
	return http.StatusOK, nil
}
func (Table1Mapper) GetTable1(nameOrId string) (int, error, interface{}) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err, nil
	}
	defer currentDb.Close()

	table1 := model.Table1{}
	table1.Field1 = nameOrId

	responseDB := currentDb.Where(&table1).Find(&table1)
	if responseDB.Error != nil {
		return http.StatusInternalServerError, responseDB.Error, nil
	}

	return http.StatusOK, nil, table1
}

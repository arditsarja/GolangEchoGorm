package mapper

import (
	"GolangEchoGorm/db"
	"GolangEchoGorm/model"
	"net/http"
)

type Table2Mapper struct {
	Table2 model.Table2
}

func (t Table2Mapper) CreateTable2() (int, error) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer currentDb.Close()
	responseDB := currentDb.Save(&t.Table2)
	if responseDB.Error != nil {
		return http.StatusInternalServerError, responseDB.Error
	}
	return http.StatusOK, nil
}

func (Table2Mapper) GetTable2(nameOrId string) (int, error, interface{}) {
	currentDb, err := db.Database()
	if err != nil {
		return http.StatusInternalServerError, err, nil
	}
	defer currentDb.Close()

	table2 := model.Table2{}

	responseDB := currentDb.Where(&table2).Find(&table2)
	if responseDB.Error != nil {
		return http.StatusInternalServerError, responseDB.Error, nil
	}

	return http.StatusOK, nil, table2
}

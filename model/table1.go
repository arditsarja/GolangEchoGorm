package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Table1 struct {
	gorm.Model
	Field1 string      `json:"field_1"`
	Field2 json.Number `json:"field_2"`
	Field3 int         `json:"field_3"`
	Field4 time.Time   `json:"field_4"`
	Field5 bool        `json:"field_5"`
	Table2 []Table2    `json:"table_2" gorm:"foreignkey:table_1_id;association_foreignkey:ID"`
	Table3 Table3      `json:"table_3"`
}

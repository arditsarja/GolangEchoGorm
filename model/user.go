package model

import (
	"GolangEchoGorm/util"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" form:"username" query:"username" gorm:"type:varchar(100)unique;not null"` // set member number to unique and not null
	FirstName string `json:"first_name" form:"first_name" query:"first_name" gorm:"type:varchar(30)"`
	LastName  string `json:"last_name" form:"last_name" query:"last_name" gorm:"type:varchar(30)"`
	Password  string `json:"password" form:"password" query:"password" gorm:"type:varchar(100);not null;"`
	Email     string `json:"email" form:"email" query:"email" gorm:"type:varchar(100);unique_index"`
	Admin     bool   `json:"admin" form:"admin" query:"admin" gorm:"type:tinyint"`
	Address   string `json:"address" form:"address" query:"address" gorm:"index:addr"` // create index with name `addr` for address
	IgnoreMe  int    `gorm:"-"`                                                        // ignore this field
}

// TableName *
func (User) TableName() string {
	return "user"
}
func (u *User) HashPassword() {
	u.Password = util.CreateHash(u.Password)
}

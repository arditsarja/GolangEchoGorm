package model

import "github.com/jinzhu/gorm"

// Token model class
type Token struct {
	gorm.Model
	Username       string
	Token          string
	DateExpiration int64
}

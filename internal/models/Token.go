package models

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	id           int    `gorm:"primaryKey"`
	Token_String string `gorm:"column:Token_String;not null;unique"`
}

func (database *Database) CheckToken(tokenStr string) error {

	var resultToken Token
	currentToken := database.Connection.Model(Token{Token_String: tokenStr}).First(&resultToken)
	return currentToken.Error
}

package models

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	id           int    `gorm:"primaryKey"`
	Token_Key    string `gorm:"column:Token_Key;not null;unique"`
	Service_Name string `gorm:"column:Service_Name;not null"`
}

func (database *Database) ValidateToken(tokenStr string) (bool, error) {
	if tokenStr == "" {
		return false, nil
	}

	//Get the count of the matching tokens in the database
	var result int64
	err := database.Connection.Model(&Token{}).Where("Token_Key = ?", tokenStr).Count(&result).Error

	//In case of an error of connecting to the database
	if err != nil {
		return false, err
	}

	//If the result is found
	if result > 0 {
		return true, nil
	}
	return false, nil
}

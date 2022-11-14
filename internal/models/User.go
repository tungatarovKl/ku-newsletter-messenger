package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id         int    `gorm:"primaryKey"`
	TelegramId int64  `gorm:"column:Telegram_Id;not null"`
	FirstName  string `gorm:"column:First_Name"`
	LastName   string `gorm:"column:Last_Name"`
	ChatId     int64  `gorm:"column:Chat_Id;not null;unique"`
}

func (database *Database) NewUser(telegramId int64, firstName, lastName string, chatId int64) error {
	//Check if user already exists
	dbFind := database.Connection.Find(&User{}, "Telegram_Id = ?", telegramId)
	if dbFind.Error != nil {
		return dbFind.Error
	}
	if dbFind.RowsAffected > 0 {
		return nil
	}

	//Create new user in database
	dbCreate := database.Connection.Create(&User{
		TelegramId: telegramId,
		FirstName:  firstName,
		LastName:   lastName,
		ChatId:     chatId,
	})
	return dbCreate.Error
}

func (database *Database) GetAllUsers() ([]User, error) {
	var users []User
	dbFind := database.Connection.Find(&users)
	return users, dbFind.Error
}

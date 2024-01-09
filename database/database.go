package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"migration/models"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=hekxo dbname=usersdb password=123456 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database")

	db = connection
	return db, nil
}

func CloseDB() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func CreateUser(user *models.User) error {
	if db == nil {
		return errors.New("database connection is nil")
	}
	return db.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := db.First(&user, id)
	return &user, result.Error
}

func UpdateUserNameByID(id uint, name string) error {
	return db.Model(&models.User{}).Where("id = ?", id).Update("name", name).Error
}

func DeleteUserByID(id uint) error {
	return db.Delete(&models.User{}, id).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

package config

import (
	"fmt"
	"praktikum/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "admin",
		DB_Password: "12345678",
		DB_Port:     "3306",
		DB_Host:     "crudgo.c2t69q0dorvz.ap-southeast-2.rds.amazonaws.com",
		DB_Name:     "crudgo",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	//DB, err = gorm.Open("mysql", "admin:12345678@tcp(crudgo.c2t69q0dorvz.ap-southeast-2.rds.amazonaws.com:3306)/crudgo")
	//if err != nil {
	//      panic(err)
	//}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, models.Book{}, models.Blog{})
}

package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

type Database struct {
	DB *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("cactus.db"), &gorm.Config{})
	// config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
	// 	username,
	// 	password,
	// 	addr,
	// 	name,
	// 	true,
	// 	//"Asia/Shanghai"),
	// 	"Local")

	// db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Printf("Database connection failed. Database name: %s", name)
	}

	// migrate
	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
}

func InitDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetDB() *gorm.DB {
	return InitDB()
}

func (db *Database) Init() {
	DB = &Database{
		DB: GetDB(),
	}
}

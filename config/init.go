package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slycreator/shop-for-me/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func OpenDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		panic("")
	}

		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",dbUser,dbPass,dbHost,dbName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{  SkipDefaultTransaction: true})

		if err != nil {
			panic("Failed to create a database connection")
		}
		db.AutoMigrate(&entity.User{},&entity.PasswordReset{})
		return db
}

func CloseDatabaseConnection(db *gorm.DB)  {
	dbSQL,err := db.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	dbSQL.Close()
}
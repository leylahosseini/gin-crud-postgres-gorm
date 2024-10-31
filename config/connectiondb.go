package config

import (
	"gin-crud-postgres-gorm/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection_DB() *gorm.DB {
	dsn := "host=localhost user=leyla password=123456 dbname=Learning-golang port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//	db.LogMode(true)
	if err != nil {
		fmt.Println("Error Connection Database")
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Connected to Database ...")

	db.AutoMigrate(&model.User{})

	return db
}

package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func main() {
	dsn := "host=localhost user=testuser password=123456 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=testuser dbname=testdb password=123456 sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		d, err := db.DB()
		if err == nil {
			d.Close()
		}
	}()

	fmt.Println("Successfully connect to new database")

	db.AutoMigrate(&User{})
}

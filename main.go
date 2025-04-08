package main

import (
	"GoApp/internal/property"
	"GoApp/routes"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	dsn := "root:@tcp(127.0.0.1:3306)/echoapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&property.Property{})

	routes.RegisterAllRoutes(e, db)
	e.Start("localhost:8080")
	// e.Logger.Fatal(e.Start(":8080"))
}

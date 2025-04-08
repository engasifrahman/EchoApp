package routes

import (
	"GoApp/internal/property"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func RegisterAllRoutes(e *echo.Echo, db *gorm.DB) {
	api := e.Group("/api/v1")

	api.GET("/heath_check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello Go!")
	})

	// Property routes
	propertyRepo := property.NewRepository(db) // Implemented somewhere
	propertyUC := property.NewUseCase(propertyRepo)
	propertyHandler := property.NewHandler(propertyUC)
	property.RegisterRoutes(api, propertyHandler)
}

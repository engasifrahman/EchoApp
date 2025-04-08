package property

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Group, h *Handler) {
    e.POST("/properties", h.CreateProperty)
    e.GET("/properties", h.ListProperties)
    e.GET("/properties:id", h.GetProperty)
    e.PUT("/properties:id", h.UpdateProperty)
    e.DELETE("/properties:id", h.DeleteProperty)
}

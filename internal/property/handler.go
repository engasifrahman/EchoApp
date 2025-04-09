package property

import (
	"net/http"

	"GoApp/internal/common"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase UseCase
}

func NewHandler(u UseCase) *Handler {
	return &Handler{usecase: u}
}

func init() {
	_ = validator.RegisterAlphaSpacesNumbersValidation(validator.Validate)
}

func (h *Handler) CreateProperty(c echo.Context) error {
	var input CreatePropertyRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	if err := validator.Validate.Struct(input); err != nil {
		// return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		formatted_error := validator.FormatValidationError(err)
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"errors": formatted_error})
	}

	prop, err := h.usecase.CreateProperty(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, prop)
}

func (h *Handler) GetProperty(c echo.Context) error {
	id := c.Param("id")
	prop, err := h.usecase.GetProperty(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "property not found"})
	}
	return c.JSON(http.StatusOK, prop)
}

func (h *Handler) ListProperties(c echo.Context) error {
	props, err := h.usecase.ListProperties(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, props)
}

func (h *Handler) UpdateProperty(c echo.Context) error {
	id := c.Param("id")
	var input UpdatePropertyRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	err := h.usecase.UpdateProperty(c.Request().Context(), id, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteProperty(c echo.Context) error {
	id := c.Param("id")
	if err := h.usecase.DeleteProperty(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

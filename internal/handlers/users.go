// users login endpoint.
package handlers

import (
	"net/http"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Login(c echo.Context) error {
	var credentials models.User
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input."})
	}

	user, err := h.Service.Login(credentials.Username, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid username or password"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Register(c echo.Context) error {
	var newUser models.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input."})
	}
	err := h.Service.Register(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to cerate user"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "User created."})
}

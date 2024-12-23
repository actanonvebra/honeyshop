// users login endpoint.
package handlers

import (
	"log"
	"net/http"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service    services.UserService
	LogService services.LogService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// login handlers
var failedAttempts = make(map[string]int)

func (h *UserHandler) Login(c echo.Context) error {
	var credentials models.User
	ip := c.RealIP()
	if err := c.Bind(&credentials); err != nil {
		log.Printf("Bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	log.Printf("Credentials received: %+v", credentials)
	user, err := h.Service.Login(credentials.Username, credentials.Password, ip)
	if err != nil {
		log.Printf("Login error for IP: %s: %v", ip, err)
		failedAttempts[ip]++
		if failedAttempts[ip] >= 5 {
			h.LogService.LogAttack("Brute Force", "Multiple failed login attempts", ip)
		}
		log.Printf("Login error: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	delete(failedAttempts, ip)
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Register(c echo.Context) error {
	var newUser models.User
	//json bind models->user.go
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// call the service layer.
	newUser, err := h.Service.Register(newUser.Username, newUser.Password, newUser.Email)
	if err != nil {
		log.Printf("Register Error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to creat user."})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})

}

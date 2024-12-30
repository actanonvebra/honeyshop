// payment endpoint.
package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/labstack/echo/v4"
)

type CheckoutHandler struct {
	Service services.CheckoutService
}

func NewCheckoutHandler(service services.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{Service: service}
}

// @Summary Checkout
// @Description Process a checkout for a user
// @Tags checkout
// @Accept json
// @Produce json
// @Param checkout body models.Checkout true "Checkout data"
// @Success 200 {string} string "Checkout completed successfully"
// @Failure 400 {string} string "Invalid Checkout"
// @Failure 500 {string} string "Checkout processing failed"
// @Router /checkout [post]
func (h *CheckoutHandler) Checkout(c echo.Context) error {
	var checkout models.Checkout
	if err := c.Bind(&checkout); err != nil {
		log.Printf("Bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Checkout"})
	}
	checkout.CreatedAt = time.Now().Format(time.RFC3339)

	// Sadece err kontrol ediyoruz, result kaldırıldı
	err := h.Service.ProcessCheckout(checkout)
	if err != nil {
		log.Printf("Checkout processing error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Checkout processing failed"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Checkout completed successfully"})
}

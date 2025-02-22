package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/labstack/echo/v4"
)

type CartHandlerInterface interface {
	CheckoutCart(c echo.Context) error
}

type CartHandler struct {
	CartService     services.CartService
	ProductService  services.ProductService
	CheckoutService services.CheckoutService
}

func NewCartHandler(cartService services.CartService, productService services.ProductService, checkoutService services.CheckoutService) *CartHandler {
	return &CartHandler{
		CartService:     cartService,
		ProductService:  productService,
		CheckoutService: checkoutService,
	}
}

// @Summary Checkout Cart
// @Description Processes the user's cart and creates a checkout record
// @Tags cart
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Success 200 {object} map[string]interface{} "Checkout completed successfully"
// @Failure 404 {string} string "Cart not found"
// @Failure 500 {string} string "Failed to process checkout"
// @Router /cart/checkout/{userID} [post]
func (h *CartHandler) CheckoutCart(c echo.Context) error {
	userID := c.Param("userID")
	log.Printf("Received userID: %s", userID)
	log.Printf("Raw userID from request: %q", c.Param("userID"))

	cart, err := h.CartService.GetCartByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "cart not found"})
	}

	var totalPrice float64
	for _, productID := range cart.Products {
		product, err := h.ProductService.GetProductByID(productID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching product details"})
		}
		totalPrice += product.Price
	}

	checkout := models.Checkout{
		ID:        "generated-id",
		UserID:    userID,
		Total:     totalPrice,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = h.CheckoutService.CreateCheckout(checkout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to clear cart"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Checkout completed successfully",
		"total":   totalPrice,
	})
}

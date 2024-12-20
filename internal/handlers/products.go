// product listing endpoint.
package handlers

import (
	"net/http"

	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.Service.FetchAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Failed to fetch products"})
	}
	return c.JSON(http.StatusOK, products)
}

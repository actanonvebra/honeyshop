// product listing endpoint.
package handlers

import (
	"log"
	"net/http"

	"github.com/actanonvebra/honeyshop/internal/helpers"
	"github.com/actanonvebra/honeyshop/internal/models"
	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service services.ProductService
	LogRepo services.LogService
}

func NewProductHandler(service services.ProductService, logService services.LogService) *ProductHandler {
	return &ProductHandler{Service: service, LogRepo: logService}
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.Service.FetchAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Failed to fetch products"})
	}
	return c.JSON(http.StatusOK, products)
}

// sqli
func (h *ProductHandler) SearchProducts(c echo.Context) error {
	search := c.QueryParam("search")
	ip := c.RealIP()

	if helpers.DetectSQLInjection(search) {
		log.Printf("Potential SQL Inejction detected: %s from IP: %s", search, ip)

		err := h.LogRepo.LogAttack("SQL Injection", search, ip)
		if err != nil {
			log.Printf("Error logging attack: %v", err)
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Invalid input detection"})

	}

	products, err := h.Service.SearchProducts(search)
	if err != nil {
		log.Printf("Error searching products: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search products"})
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) AddProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		log.Println("c.Bind product error", err)
	}
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 || product.Category == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required parameter"})
	}
	err := h.Service.AddProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add product"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Product added successfully"})

}

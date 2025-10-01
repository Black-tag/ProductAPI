package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Black-tag/productAPI/internal/database"
	"github.com/Black-tag/productAPI/internal/logger"
	"github.com/Black-tag/productAPI/internal/models"
	"github.com/google/uuid"
)
// @Summary Create Products
// @Description Existing users can create aproducts
// @Tags products
// @Accept json
// @Produce json
// @Param request body models.ProductCreationRequest true "Product creation data"
// @Success 201 {object} models.ProductCreationResponse
// @Failure 400 {object} string "Bad Request - Invalid input"
// @Failure 404 {object} string "Not Found - Resource doesn't exist"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/product [post]
// @Security BearerAuth
func (cfg *APIConfig) ProductCreationHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("entered Product creation handler")
	userIDValue := r.Context().Value("userID")
	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		http.Error(w, "userID not in context", http.StatusInternalServerError)
		return
	}

	var req models.ProductCreationRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	product, err := cfg.DB.CreateProductsFromRequest(r.Context(), database.CreateProductsFromRequestParams{
		Name:     req.Name,
		Price:    fmt.Sprintf("%.2f", req.Price),
		PostedBy: userID,
	})
	if err != nil {
		http.Error(w, "databse operation to create product failed", http.StatusInternalServerError)
		return
	}

	data := models.ProductCreationResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		PostedBy:  product.PostedBy,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode products", http.StatusInternalServerError)
		return
	}

}
// @Summary Get existing  products
// @Description  users can get all existing Products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} database.Product
// @Failure 400 {object} string "Bad Request - Invalid input"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/product [get]
// @Security BearerAuth
func (cfg *APIConfig) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("entered get products handler")

	products, err := cfg.DB.GetAllProducts(r.Context())
	if err != nil {
		http.Error(w, "databse opertaion to get products failed", http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		http.Error(w, "no products in databse", http.StatusNoContent)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "failed to encode products", http.StatusInternalServerError)
		return
	}
}


// @Summary Delete an existing  product
// @Description admin can delete product using their id
// @Tags products
// @Accept json
// @Produce json
// @Success 204 {string} string "No content"
// @Failure 400 {object} string "Bad Request - Invalid input"
// @Failure 401 {object} string "Unauthorized - Missing/invalid credentials"
// @Failure 403 {object} string "Forbidden - Insufficient permissions"
// @Failure 500 {object} string "Internal Server Error"
// @Param productID path string true "ProductID" 
// @Router /api/v1/product/{productID} [delete]
// @Security BearerAuth
func (cfg *APIConfig) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("entered product deletion handler")

	productIdStr := r.PathValue("productID")
	userIDval := r.Context().Value("userID")
	userID, ok := userIDval.(uuid.UUID)
	if !ok {
		http.Error(w, "userID not in context", http.StatusUnauthorized)
		return
	}

	userRole := r.Context().Value("role")

	productID, err := uuid.Parse(productIdStr)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return

	}

	product, err := cfg.DB.GetProductByID(r.Context(), productID)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if userID != product.PostedBy && userRole != "admin" {
		http.Error(w, "unauthorised for deleting", http.StatusUnauthorized)
		return

	}

	err = cfg.DB.DeleteProductByID(r.Context(), productID)
	if err != nil {
		http.Error(w, "databse deletion failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// @Summary Update an existing  product
// @Description Existing users can update their product
// @Tags products
// @Accept json
// @Produce json
// @Param bugid path string true "productID" example:"87f0ea02-7b24-41bd-8418-0831a019fc87"
// @Param request body models.UpdateProductRequest true "product updation data"
// @Success 200 {object} models.UpdatedProductResponse
// @Failure 400 {object} string "Bad Request - Invalid input"
// @Failure 401 {object} string "Unauthorized - Missing/invalid credentials"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/product/{productID} [put]
// @Security BearerAuth
func (cfg *APIConfig) UpdateProductsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("entered update products Hnadler")
	userIDValue := r.Context().Value("userID")

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		http.Error(w, "userID not in context ", http.StatusUnauthorized)
		return
	}
	userrole := r.Context().Value("role")

	productIDStr := r.PathValue("productID")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		http.Error(w, "cannot parse product id into uuid", http.StatusBadRequest)
		return
	}
	var req models.UpdateProductRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "cannot decode request", http.StatusInternalServerError)
		return

	}
	product, err := cfg.DB.GetProductByID(r.Context(), productID)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if userID != product.PostedBy && userrole != "admin" {
		http.Error(w, "unauthorized for edit product", http.StatusUnauthorized)
		return

	}
	updatedProduct, err := cfg.DB.UpdateProduct(r.Context(), database.UpdateProductParams{
		ID:    productID,
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		http.Error(w, "databse operation failed", http.StatusInternalServerError)
		return
	}

	respPayload := models.UpdatedProductResponse{
		ID:        updatedProduct.ID,
		Name:      updatedProduct.Name,
		Price:     updatedProduct.Price,
		CreatedAt: updatedProduct.CreatedAt,
		UpdatedAt: updatedProduct.UpdatedAt,
		PostedBy:  updatedProduct.PostedBy,
	}

	if err := json.NewEncoder(w).Encode(respPayload); err != nil {
		http.Error(w, "failed to encode products", http.StatusInternalServerError)
		return
	}
}

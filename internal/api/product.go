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

func (cfg *APIConfig) ProductCreationHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("entered Product creation handler")
	userIDValue := r.Context().Value("userID")
	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		http.Error(w, "userID not in context", http.StatusInternalServerError)
		return
	}

	var req models.ProductCreationrequest

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
		CreatedAt: product.CreatedAt.Time,
		UpdatedAt: product.UpdatedAt.Time,
		PostedBy:  product.PostedBy,
	}
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "cannot marshal struct into json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}



func (cfg *APIConfig) GetProductsHandler (w http.ResponseWriter, r *http.Request) {
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

func (cfg *APIConfig) DeleteProducthandler (w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("entered product deletion handler")
	productIdStr := r.PathValue("productID")
	productID, err := uuid.Parse(productIdStr)
	if err != nil {
		http.Error(w, "cannot parse str into uuid", http.StatusInternalServerError)
		return

	}

	err = cfg.DB.DeleteProductByID(r.Context(), productID)
	if err != nil {
		http.Error(w, "databse deletion failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)




}

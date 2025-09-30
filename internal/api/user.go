package api

import (
	"encoding/json"
	"time"

	"net/http"

	"github.com/Black-tag/productAPI/internal/logger"
	"github.com/Black-tag/productAPI/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	
	logger.Log.Info("entered user creation handler")

	w.Header().Set("Content-Type", "application/json")

	var req models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w,"bad request format", http.StatusBadRequest)
		return
	}
	logger.Log.Info("captured request",
	zap.String("email_in_request", req.Email),
	zap.String("password_in_request", req.Password),
	)

	payload := models.UserResponse{
		Id: uuid.New(),
		Email: req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	}
	logger.Log.Info("user_response_payload",
		zap.String("userID", payload.Id.String()),
		zap.String("user_emai", payload.Email),
		zap.Time("user_Created_at", payload.CreatedAt),
		zap.Time("user_updated_at", payload.UpdatedAt),

	)
	resp, err := json.Marshal(payload)
	if err != nil {
		logger.Log.Error("cannot marshal payload to response in json")
		http.Error(w, "json marshalling error", http.StatusInternalServerError)
		return

	}
	
	w.Write(resp)

}

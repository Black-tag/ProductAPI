package api

import (
	"encoding/json"

	"net/http"

	"github.com/Black-tag/productAPI/internal/database"
	"github.com/Black-tag/productAPI/internal/logger"
	"github.com/Black-tag/productAPI/internal/models"
	"github.com/Black-tag/productAPI/internal/utils"

	"go.uber.org/zap"
)

func (cfg *APIConfig)CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	
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

	hashdepassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "cannot hash password", http.StatusInternalServerError)
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Email: req.Email,
		Hashedpassword: hashdepassword,
	})
	if err != nil {
		logger.Log.Error("cannot create user in databse")
		http.Error(w, "databse operation failed", http.StatusInternalServerError)
		return
	}
	respPayload := database.User{
		ID: user.ID,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role: user.Role,

	}

	
	logger.Log.Info("user_response_payload",
		zap.String("userID", user.ID.String()),
		zap.String("user_emai", user.Email),
		zap.Time("user_Created_at", user.CreatedAt.Time),
		zap.Time("user_updated_at", user.UpdatedAt.Time),

	)
	resp, err := json.Marshal(respPayload)
	if err != nil {
		logger.Log.Error("cannot marshal payload to response in json")
		http.Error(w, "json marshalling error", http.StatusInternalServerError)
		return

	}
	
	w.Write(resp)

}

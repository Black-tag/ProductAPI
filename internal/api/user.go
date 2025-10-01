package api

import (
	"database/sql"
	"encoding/json"
	"time"

	"net/http"

	"github.com/Black-tag/productAPI/internal/database"
	"github.com/Black-tag/productAPI/internal/logger"
	"github.com/Black-tag/productAPI/internal/models"
	"github.com/Black-tag/productAPI/internal/utils"

	"go.uber.org/zap"
)

// @Summary Creates a new  user
// @Description Creates user with Email and Password
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.UserRequest true "User creation data"
// @Success 201 {object} database.User
// @Failure 400 {object} string "Bad Request - Invalid input"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/users [post]
// @Security BearerAuth
func (cfg *APIConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	logger.Log.Info("entered user creation handler")

	w.Header().Set("Content-Type", "application/json")

	var req models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request format", http.StatusBadRequest)
		return
	}
	logger.Log.Info("captured request",
		zap.String("email_in_request", req.Email),
	)

	hashdepassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "cannot hash password", http.StatusInternalServerError)
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Email:          req.Email,
		Hashedpassword: hashdepassword,
	})
	if err != nil {
		logger.Log.Error("cannot create user in databse")
		http.Error(w, "databse operation failed", http.StatusInternalServerError)
		return
	}
	respPayload := database.User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role:      user.Role,
	}

	logger.Log.Info("user_response_payload",
		zap.String("userID", user.ID.String()),
		zap.String("user_emai", user.Email),
		zap.Time("user_Created_at", user.CreatedAt),
		zap.Time("user_updated_at", user.UpdatedAt),
	)
	resp, err := json.Marshal(respPayload)
	if err != nil {
		logger.Log.Error("cannot marshal payload to response in json")
		http.Error(w, "json marshalling error", http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}


// @Summary Login an existing  user
// @Description Existing users can login using email and password
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "user login data"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} string "Bad Request - Invalid input"
// @Failure 401 {object} string "Invalid credentials"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/login [post]
// @Security BearerAuth
func (cfg *APIConfig) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := cfg.DB.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		http.Error(w, "user does not exists", http.StatusNotFound)
		return
	}
	if err := utils.CheckPasswordAndHash(req.Password, user.Hashedpassword); err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := utils.MakeJWT(user.ID, cfg.SECRET, time.Hour)
	if err != nil {
		http.Error(w, "cannot create jwt", http.StatusInternalServerError)
		return
	}
	refreshToken, err := utils.MakeRefreshToken()
	if err != nil {
		http.Error(w, "cannot create refresh token", http.StatusInternalServerError)
		return
	}
	refreshExpiresAt := time.Now().Add(30 * 24 * time.Hour)

	err = cfg.DB.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: refreshExpiresAt,
		RevokedAt: sql.NullTime{},
	})
	if err != nil {
		http.Error(w, "cannot create refresh toke", http.StatusInternalServerError)
		return
	}
	respPayload := models.LoginResponse{
		ID:           user.ID,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Role:         user.Role,
		Token:        token,
		RefreshToken: refreshToken,
	}
	resp, err := json.Marshal(respPayload)
	if err != nil {
		http.Error(w, "cannot marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

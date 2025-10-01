package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id        uuid.UUID `json:"userID"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"cretaed_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Role         string    `json:"role"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

type ProductCreationrequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductCreationResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PostedBy  uuid.UUID `json:"posted_by"`
}

type UpdateProductRequest struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type UpdatedProductResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PostedBy  uuid.UUID `json:"posted_by"`
}

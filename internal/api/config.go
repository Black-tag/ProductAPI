package api

import (
	"github.com/Black-tag/productAPI/internal/database"
)

type APIConfig struct {
	DB     *database.Queries
	SECRET string
}

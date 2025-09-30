package main

import (
	"database/sql"
	"net/http"
	"os"
	_ "github.com/lib/pq"

	"github.com/Black-tag/productAPI/internal/api"
	"github.com/Black-tag/productAPI/internal/database"

	"github.com/Black-tag/productAPI/internal/logger"
)

func main() {

	logger.Init()
	defer logger.Log.Sync()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		logger.Log.Fatal("DB_URL env variable not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	dbQueries := database.New(db)

	cfg := api.APIConfig{
		DB: dbQueries,
	}


	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/users", cfg.CreateUserHandler)



	logger.Log.Info("server starting on 8090")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		logger.Log.Fatal("server not started")
		return
	}
	
}

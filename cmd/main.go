package main

import (
	"net/http"

	"github.com/Black-tag/productAPI/internal/api"
	"github.com/Black-tag/productAPI/internal/logger"
	
)

func main() {

	logger.Init()
	defer logger.Log.Sync()


	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/users", api.CreateUserHandler)



	logger.Log.Info("server starting on 8090")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		logger.Log.Fatal("server not started")
		return
	}
	
}

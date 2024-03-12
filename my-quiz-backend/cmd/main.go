package main

import (
	"fmt"
	"my-quiz-backend/api/v1"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	api.SetupRoutes(logger)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

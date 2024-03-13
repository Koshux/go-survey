package main

import (
	"fmt"
	"my-quiz-backend/api/v1"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	router := mux.NewRouter()
	var quizApi = router.PathPrefix("/api").Subrouter()

	api.SetupRoutes(logger, quizApi)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

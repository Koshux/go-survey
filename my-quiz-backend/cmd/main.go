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
	router := setupMuxRouter(logger)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func setupMuxRouter(logger *zap.Logger) *mux.Router {
	router := mux.NewRouter()
	// Handle all preflight request
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
	})

	router.StrictSlash(true)
	var quizApi = router.PathPrefix("/api").Subrouter()
	quizApi.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	api.SetupRoutes(logger, quizApi)

	return router
}

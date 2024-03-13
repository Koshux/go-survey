package api

import (
	"my-quiz-backend/internal/quiz"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func SetupRoutes(logger *zap.Logger, quizApi *mux.Router) {
	var v1Api = quizApi.PathPrefix("/v1/quiz").Subrouter()

	setupMiddleware(logger, v1Api)
	setupHandlers(logger, v1Api)
}

func setupMiddleware(logger *zap.Logger, api *mux.Router) {
	api.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Request received", zap.String("path", r.URL.Path))
			next.ServeHTTP(w, r)
		})
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupHandlers(logger *zap.Logger, api *mux.Router) {
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.WriteHeader(http.StatusNotFound)
	})

	api.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		quiz.GetQuestions(w, logger)
	}).Methods("GET")

	api.HandleFunc("/answers", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		quiz.SubmitQuizAnswers(w, r, logger)
	}).Methods("POST")

	api.HandleFunc("/answers", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		quiz.GetQuizResults(w, logger)
	}).Methods("GET")
}

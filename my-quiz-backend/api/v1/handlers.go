package api

import (
	"my-quiz-backend/internal/quiz"
	"net/http"

	"go.uber.org/zap"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func SetupRoutes(logger *zap.Logger) {
	http.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			quiz.GetQuestions(w, logger)
		}
	})

	http.HandleFunc("/answers", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		switch r.Method {
		case "GET":
			quiz.GetQuizResults(w, logger)
		case "POST":
			quiz.SubmitQuizAnswers(w, r, logger)
		}
	})
}

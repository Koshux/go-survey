package quiz

import (
	"encoding/json"
	"my-quiz-backend/models"
	"my-quiz-backend/store"
	"net/http"

	"go.uber.org/zap"
)

func GetQuestions(w http.ResponseWriter, logger *zap.Logger) {
	questions := store.GetQuestions()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)

	logger.Info("Retrieve quiz questions", zap.String("status", "completed"))
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	logger.Info("Submitting answers", zap.String("status", "started"))

	var submittedAnswer models.SubmittedAnswer
	err := json.NewDecoder(r.Body).Decode(&submittedAnswer)

	if err != nil {
		errorAnswer := zap.Any("errorAnswer", submittedAnswer)
		logger.Error("Failed to submit the answer", errorAnswer)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logSubmittedAnswer := zap.Any("logSubmittedAnswer", submittedAnswer)
	logger.Info("Pending submitted answer", logSubmittedAnswer)

	store.AddSubmittedAnswer(submittedAnswer)

	logger.Info("Submitted answers completed", zap.String("status", "success"))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answers submitted successfully",
	})
}

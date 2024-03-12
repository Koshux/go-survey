package quiz

import (
	"encoding/json"
	"my-quiz-backend/models"
	"my-quiz-backend/store"
	"net/http"

	"github.com/google/uuid"
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
		logger.Error("Failed to submit the answer", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logSubmittedAnswer := zap.Any("logSubmittedAnswer", submittedAnswer)
	logger.Info("Pending submitted answer", logSubmittedAnswer)

	score := calculateScore(submittedAnswer)
	store.AddQuizResult((models.QuizResult{
		UserID: uuid.New(),
		Score:  score,
	}))

	response := map[string]interface{}{
		"message":    "Answers submitted successfully",
		"score":      score,
		"percentile": store.CalculatePercentile(submittedAnswer.UserID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	logger.Info("Submitted answers completed", zap.String("status", "success"))
}

func calculateScore(submittedAnswer models.SubmittedAnswer) int {
	score := 0

	for questionID, userAnswer := range submittedAnswer.Answers {
		correctAnswer, exists := store.GetCorrectAnswer(questionID)
		if exists && correctAnswer == userAnswer {
			score++
		}
	}

	return score
}

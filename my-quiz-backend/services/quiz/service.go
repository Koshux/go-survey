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

	logger.Debug("Retrieve quiz questions", zap.Any("payload", questions))
	logger.Info("Retrieve quiz questions", zap.String("status", "completed"))
}

func SubmitQuizAnswers(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	logger.Info("Submitting answers", zap.String("status", "started"))

	var submittedAnswer models.SubmittedAnswer
	err := json.NewDecoder(r.Body).Decode(&submittedAnswer)

	if err != nil {
		logger.Error("Failed to submit the answer", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logSubmittedAnswer := zap.Any("logSubmittedAnswer", submittedAnswer)
	logger.Debug("Pending submitted answer", logSubmittedAnswer)

	score := calculateScore(submittedAnswer)
	userID := uuid.New()

	quizResult := models.QuizResult{
		UserID:     userID,
		Score:      score,
		Percentile: 0,
	}

	// First add the submitted answer to the store then calculate the percentile
	// because the quiz result must be added to the store before calculating the
	// percentile. Otherwise the user's score will not be included in the scores
	// used to calculate the percentile.
	store.AddQuizResult(quizResult)
	quizResult.Percentile = store.CalculatePercentile(logger, userID, score)

	response := map[string]interface{}{
		"message":    "Answers submitted successfully",
		"score":      score,
		"percentile": quizResult.Percentile,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	logger.Debug("Submitted answers completed", zap.Any("payload", response))
	logger.Info("Submitted answers completed", zap.String("status", "success"))
}

func GetQuizResults(w http.ResponseWriter, logger *zap.Logger) {
	quizResults := store.GetQuizResults()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quizResults)

	logger.Debug("Retrieve submitted answers", zap.Any("payload", quizResults))
	logger.Info("Retrieve submitted answers", zap.String("status", "completed"))
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

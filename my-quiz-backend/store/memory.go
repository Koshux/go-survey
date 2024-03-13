package store

import (
	"fmt"
	"my-quiz-backend/models"
	"sort"

	"github.com/google/uuid"
)

var questions = []models.Question{
	{
		ID:      "1",
		Text:    "What is the capital of France?",
		Options: []string{"New York", "London", "Paris", "Tokyo"},
		Answer:  2, // Paris
	},
	{
		ID:      "2",
		Text:    "Which planet is known as the Red Planet?",
		Options: []string{"Earth", "Mars", "Jupiter", "Saturn"},
		Answer:  1, // Mars
	},
	{
		ID:      "3",
		Text:    "What is the largest ocean on Earth?",
		Options: []string{"Atlantic Ocean", "Indian Ocean", "Arctic Ocean", "Pacific Ocean"},
		Answer:  3, // Pacific Ocean
	},
	{
		ID:      "4",
		Text:    "In which year did the Titanic sink?",
		Options: []string{"1912", "1922", "1905", "1898"},
		Answer:  0, // 1912
	},
	{
		ID:      "5",
		Text:    "Who wrote 'Hamlet'?",
		Options: []string{"Leo Tolstoy", "William Shakespeare", "Charles Dickens", "Mark Twain"},
		Answer:  1, // William Shakespeare
	},
	{
		ID:      "6",
		Text:    "What element does 'O' represent on the periodic table?",
		Options: []string{"Gold", "Oxygen", "Silver", "Iron"},
		Answer:  1, // Oxygen
	},
	{
		ID:      "7",
		Text:    "Who is known as the father of Geometry?",
		Options: []string{"Aristotle", "Euclid", "Pythagoras", "Kepler"},
		Answer:  1, // Euclid
	},
	{
		ID:      "8",
		Text:    "What is the capital of Japan?",
		Options: []string{"Seoul", "Beijing", "Tokyo", "Bangkok"},
		Answer:  2, // Tokyo
	},
	{
		ID:      "9",
		Text:    "Which is the largest planet in our solar system?",
		Options: []string{"Earth", "Jupiter", "Saturn", "Mars"},
		Answer:  1, // Jupiter
	},
	{
		ID:      "10",
		Text:    "Who wrote the novel '1984'?",
		Options: []string{"George Orwell", "Aldous Huxley", "Ray Bradbury", "H.G. Wells"},
		Answer:  0, // George Orwell
	},
}
var quizResults []models.QuizResult
var submittedAnswers []models.SubmittedAnswer

func CalculatePercentile(userID uuid.UUID) float64 {
	var userScore int
	var scores []int

	for _, result := range quizResults {
		scores = append(scores, result.Score)
		fmt.Println(">> CalculatePercentile - result.Score:", result.Score)
		if result.UserID == userID {
			userScore = result.Score
		}
	}
	fmt.Println(">> CalculatePercentile - userScore:", userScore)

	sort.Ints(scores)

	// Find the rank of the user's score
	rank := float64(len(scores))
	for i, score := range scores {
		if score == userScore {
			rank = float64(i)
			break
		}
	}

	fmt.Println(">> CalculatePercentile - Rank:", rank)
	fmt.Println(">> CalculatePercentile - scores:", scores)
	fmt.Println(">> CalculatePercentile - len(scores):", len(scores))
	percentile := (rank / float64(len(scores))) * 100
	return percentile
}

func GetCorrectAnswer(questionID string) (string, bool) {
	for _, q := range questions {
		if q.ID == questionID {
			return q.Options[q.Answer], true
		}
	}

	return "", false
}

func GetQuestions() []models.Question {
	return questions
}

func AddSubmittedAnswer(answer models.SubmittedAnswer) {
	submittedAnswers = append(submittedAnswers, answer)
}

func GetSubmittedAnswers() []models.SubmittedAnswer {
	return submittedAnswers
}

func AddQuizResult(result models.QuizResult) {
	quizResults = append(quizResults, result)
}

func GetQuizResults() []models.QuizResult {
	return quizResults
}

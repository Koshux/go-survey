package store

import (
	"math"
	"my-quiz-backend/models"
	"sort"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var questions = []models.Question{
	{
		ID: "1",
		Text: map[string]string{
			"default": "What is the capital of France?",
			"de":      "Was ist die Hauptstadt von Frankreich?",
			"pt":      "Qual é a capital da França?",
			"sv":      "Vad är huvudstaden i Frankrike?",
		},
		Options: []string{"New York", "London", "Paris", "Tokyo"},
		Answer:  2, // Paris
	},
	{
		ID: "2",
		Text: map[string]string{
			"default": "Which planet is known as the Red Planet?",
			"de":      "Welcher Planet ist als der Rote Planet bekannt?",
			"pt":      "Qual planeta é conhecido como o Planeta Vermelho?",
			"sv":      "Vilken planet är känd som den röda planeten?",
		},
		Options: []string{"Earth", "Mars", "Jupiter", "Saturn"},
		Answer:  1, // Mars
	},
	{
		ID: "3",
		Text: map[string]string{
			"default": "What is the largest ocean on Earth?",
			"de":      "Was ist der größte Ozean der Erde?",
			"pt":      "Qual é o maior oceano da Terra?",
			"sv":      "Vilken är den största oceanen på jorden?",
		},
		Options: []string{"Atlantic Ocean", "Indian Ocean", "Arctic Ocean", "Pacific Ocean"},
		Answer:  3, // Pacific Ocean
	},
	{
		ID: "4",
		Text: map[string]string{
			"default": "In which year did the Titanic sink?",
			"de":      "In welchem Jahr sank die Titanic?",
			"pt":      "Em que ano o Titanic afundou?",
			"sv":      "Vilket år sjönk Titanic?",
		},
		Options: []string{"1912", "1922", "1905", "1898"},
		Answer:  0, // 1912
	},
	{
		ID: "5",
		Text: map[string]string{
			"default": "Who wrote 'Hamlet'?",
			"de":      "Wer schrieb 'Hamlet'?",
			"pt":      "Quem escreveu 'Hamlet'?",
			"sv":      "Vem skrev 'Hamlet'?",
		},
		Options: []string{"Leo Tolstoy", "William Shakespeare", "Charles Dickens", "Mark Twain"},
		Answer:  1, // William Shakespeare
	},
	{
		ID: "6",
		Text: map[string]string{
			"default": "What element does 'O' represent on the periodic table?",
			"de":      "Welches Element repräsentiert 'O' im Periodensystem?",
			"pt":      "Qual elemento 'O' representa na tabela periódica",
			"sv":      "Vilket element representerar 'O' på det periodiska systemet",
		},
		Options: []string{"Gold", "Oxygen", "Silver", "Iron"},
		Answer:  1, // Oxygen
	},
	{
		ID: "7",
		Text: map[string]string{
			"default": "Who is known as the father of Geometry?",
			"de":      "Wer ist als der Vater der Geometrie bekannt?",
			"pt":      "Quem é conhecido como o pai da Geometria?",
			"sv":      "Vem är känd som geometriens far?",
		},
		Options: []string{"Aristotle", "Euclid", "Pythagoras", "Kepler"},
		Answer:  1, // Euclid
	},
	{
		ID:      "8",
		Options: []string{"Seoul", "Beijing", "Tokyo", "Bangkok"},
		Text: map[string]string{
			"default": "What is the capital of Japan?",
			"de":      "Was ist die Hauptstadt von Japan?",
			"pt":      "Qual é a capital do Japão?",
			"sv":      "Vad är huvudstaden i Japan?",
		},
		Answer: 2, // Tokyo
	},
	{
		ID: "9",
		Text: map[string]string{
			"default": "Which is the largest planet in our solar system?",
			"de":      "Welcher ist der größte Planet in unserem Sonnensystem?",
			"pt":      "Qual é o maior planeta em nosso sistema solar?",
			"sv":      "Vilken är den största planeten i vårt solsystem?",
		},
		Options: []string{"Earth", "Jupiter", "Saturn", "Mars"},
		Answer:  1, // Jupiter
	},
	{
		ID: "10",
		Text: map[string]string{
			"default": "Who wrote the novel '1984'?",
			"de":      "Wer schrieb den Roman '1984'?",
			"pt":      "Quem escreveu o romance '1984'?",
			"sv":      "Vem skrev romanen '1984'?",
		},
		Options: []string{"George Orwell", "Aldous Huxley", "Ray Bradbury", "H.G. Wells"},
		Answer:  0, // George Orwell
	},
}
var quizResults []models.QuizResult
var submittedAnswers []models.SubmittedAnswer

func CalculatePercentile(logger *zap.Logger, userID uuid.UUID, score int) float64 {
	if userID == uuid.Nil {
		logger.Error("Invalid userID", zap.String("userID", userID.String()))
		return 0
	}

	var userScore int
	var scores []int
	var belowCount int // The number of scores below the user's score

	if len(quizResults) == 0 {
		scores = append(scores, score)
		userScore = score
	} else {
		for _, result := range quizResults {
			scores = append(scores, result.Score)
			if result.UserID == userID {
				userScore = result.Score
			}
		}
	}

	sort.Ints(scores)

	// Increment the belowCount for each score that is less than the user's score
	for _, score := range scores {
		if score < userScore {
			belowCount++
		}
	}

	// If there are less than 7 scores, the percentile cannot be calculated
	// since the belowCount will be 0 and results in a division by zero. In this
	// case, 999 is used to indicate an error.
	N := len(scores)
	if N <= 7 {
		return 999
	}

	percentile := (float64(belowCount) / float64(N)) * 100
	return roundToTwoDecimalPlaces(percentile)
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
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

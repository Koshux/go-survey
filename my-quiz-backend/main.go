package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			getQuestions(w, r)
		} else if r.Method == "POST" {
			postAnswers(w, r)
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Question struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	Options []string `json:"options"`
	Answer  int `json:"answer"` // index of the correct option
}

var questions = []Question{
	{
		ID:			"1",
		Text:		"What is the capital of France?",
		Options:	[]string{"New York", "London", "Paris", "Tokyo"},
		Answer:		2, // Paris
	},
	{
		ID:			"2",
		Text:		"Which planet is known as the Red Planet?",
		Options:	[]string{"Earth", "Mars", "Jupiter", "Saturn"},
		Answer:		1, // Mars
	},
	{
		ID:			"3",
		Text:		"What is the largest ocean on Earth?",
		Options:	[]string{"Atlantic Ocean", "Indian Ocean", "Arctic Ocean", "Pacific Ocean"},
		Answer:		3, // Pacific Ocean
	},
	{
		ID:			"4",
		Text:		"In which year did the Titanic sink?",
		Options:	[]string{"1912", "1922", "1905", "1898"},
		Answer:		0, // 1912
	},
	{
		ID:			"5",
		Text:		"Who wrote 'Hamlet'?",
		Options:	[]string{"Leo Tolstoy", "William Shakespeare", "Charles Dickens", "Mark Twain"},
		Answer:		1, // William Shakespeare
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

func getQuestions(w http.ResponseWriter, r *http.Request) {
	// Return the questions to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}

func postAnswers(w http.ResponseWriter, r *http.Request) {
	// Process submitted answers and return the result

}

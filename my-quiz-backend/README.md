# Quiz App Backend

This document outlines the backend setup for a simple quiz application built with Go. The backend serves quiz questions, processes quiz submissions, and calculates percentile ranks for quiz takers, with a nuanced approach to evaluating performances through fuzzy logic.

## Overview

The backend is designed to:
- Serve a set of multilingual quiz questions.
- Accept quiz submissions and calculate scores.
- Calculate and return a user's percentile rank and performance category based on their score relative to all submissions.

## Technology Stack

- **Language**: Go
- **Dependencies**:
  - Standard library for core functionalities.
  - `github.com/gorilla/mux` for routing and request handling.
  - `go.uber.org/zap` for structured logging.
  - `github.com/google/uuid` for generating unique identifiers for quiz submissions.

## Setup and Running

### Prerequisites

Ensure you have Go installed on your system. Download it from [the official Go website](https://golang.org/dl/).

### Running the Server

1. Navigate to the project directory.
2. Run the server using the Go command:
   ```bash
   go run ./cmd/main.go
   ```

3. The server will start and listen on `http://localhost:8080`.

## API Endpoints

- **GET `/questions`**: Fetches a list of multilingual quiz questions.
- **POST `/answers`**: Submits answers for the quiz. This endpoint calculates the user's score, their percentile rank, returns this information, and the performance category in the response.
- **GET `/results`**: Fetches a list of all quiz submissions, including users' scores, percentile ranks, and performance categories.

### Questions Endpoint

Returns a JSON array of multilingual quiz questions, including question text, options, and the index of the correct answer.

Example request:
```plaintext
GET /questions HTTP/1.1
Host: localhost:8080
```

Example response:
```json
[
  {
    "id": "1",
    "text": "What is the capital of France?",
    "options": ["New York", "London", "Paris", "Tokyo"],
    "answer": 2
  }
]
```

### Answers Endpoint

Accepts a JSON object containing the user's answers and returns the user's score, percentile rank, and performance category.

Example request:
```plaintext
POST /answers HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "userId": "unique-user-id",
  "answers": {
    "1": "Paris",
    "2": "Mars"
  }
}
```

Example response:
```json
{
  "userID": "unique-user-id",
  "score": 2,
  "percentile": 75,
  "category": "Good to Excellent"
}
```

## Understanding the Percentile Rank and Performance Category

- The **percentile rank** indicates a user's performance relative to all quiz takers. For example, a percentile rank of 75 means the user scored better than approximately 75% of participants.
- The **performance category** provides nuanced feedback, categorizing performances into "Needs Improvement to Average", "Average to Good", "Good to Excellent", based on fuzzy logic evaluation of the scores.

## Development Notes

- The server uses in-memory storage for simplicity. Consider a database for production environments.
- CORS is enabled for all origins. Adjust according to your security requirements.
- Performance categories are determined using fuzzy logic to provide nuanced feedback beyond simple numerical scores.

## Future Enhancements

- Persistent database integration for questions and submissions.
- User authentication for personalized quiz experiences.
- API expansion for dynamic quiz management.

## License

[MIT](LICENSE)

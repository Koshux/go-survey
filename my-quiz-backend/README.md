# Quiz App Backend

This document outlines the backend setup for a simple quiz application built with Go. The backend serves quiz questions and processes quiz submissions, including calculating percentile ranks for quiz takers.

## Overview

The backend is designed to:
- Serve a set of quiz questions.
- Accept quiz submissions and calculate scores.
- Calculate and return a user's percentile rank based on their score relative to all submissions.

## Technology Stack

- **Language**: Go
- **Dependencies**:
  - Standard library for core functionalities.
  - `github.com/gorilla/mux` for routing and request handling.
  - `go.uber.org/zap` for structured logging.
  - `github.com/google/uuid` for generating unique identifiers for quiz submissions.

## Setup and Running

### Prerequisites

Ensure you have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).

### Running the Server

1. Navigate to the project directory.
2. Run the server using the Go command:
   ```bash
   go run ./cmd/main.go
   ```

3. The server will start and listen on `http://localhost:8080`.

## API Endpoints

- **GET `/questions`**: Fetches a list of quiz questions.
- **POST `/answers`**: Submits answers for the quiz. This endpoint calculates the user's score, their percentile rank, and returns this information in the response.
- **GET `/answers`**: Fetches a list of all quiz submissions which includes the user's score and percentile rank.

### Questions Endpoint

Returns a JSON array of quiz questions. Each question includes an `id`, `text`, options, and the index of the correct answer.

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

Accepts a JSON object containing the user's answers and returns the user's score and percentile rank.

Example request:
```plaintext
POST /answers HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

```json
{
  "userId": "unique-user-id",
  "answers": {
    "1": "Paris",
    "2": "Mars",
    "3": "Pacific Ocean",
    // More questionId: answer pairs
  }
}
```


Example response:

```json
{
  "userID": "unique-user-id",
  "score": 2,
  "percentile": 75
}
```

### Understanding the Percentile Rank

The percentile rank indicates your performance relative to all quiz takers. For example, a percentile rank of 87.5 means you scored better than approximately 87.5% of participants. This feedback helps gauge where you stand and encourages improvement.

## Development Notes

- The server uses in-memory storage for simplicity. For a production environment, consider integrating a persistent database.
- Cross-Origin Resource Sharing (CORS) is enabled for all origins. Adjust the CORS settings according to your security requirements.
- The percentile calculation assumes a significant number of submissions for accuracy. For early-stage testing, special handling may adjust feedback for users.

### Important Notes

- The actual answer indices (`"answer": 2`) are not included in the response to the `/questions` endpoint. This information is for documentation purposes only to illustrate the structure of a question.
- While developing or testing, you might encounter a special percentile value indicating insufficient data for a meaningful calculation. This mechanism is in place to ensure the accuracy and reliability of percentile feedback as more users participate.

## Future Enhancements

- Integrate a database for persistent storage of questions and submissions.
- Implement user authentication and session management.
- Expand the API to support dynamic creation and management of quiz questions.

## License

[MIT](LICENSE)

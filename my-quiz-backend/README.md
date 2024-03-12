# Quiz App Backend

This document outlines the backend setup for a simple quiz application built with Go. The backend serves quiz questions only for the time being.

## Overview
The backend is designed to:
- Serve a set of quiz questions.

## Technology Stack
- **Language**: Go
- **Dependencies**: None, standard library only

## Setup and Running

### Prerequisites
Ensure you have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).

### Running the Server
1. Navigate to the project directory.
2. Run the server using the Go command:
   ```bash
   go run main.go
   ```

3. The server will start and listen on `http://localhost:8080`.

## API Endpoints
- **GET `/questions`**: Fetches a list of quiz questions.

### Questions Endpoint
Returns a JSON array of quiz questions. Each question includes an `id`, `text`, options, and the index of the correct answer.

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

## Development Notes

- The server uses in-memory storage for simplicity. For a production environment, consider integrating a persistent database.
- Cross-Origin Resource Sharing (CORS) is enabled for all origins. Adjust the CORS settings according to your security requirements.

## License
[MIT](LICENSE)

# Quiz App Frontend
This frontend for the quiz app is built using Vue 3 with Vite. It utilizes Pinia for state management and SurveyJS for rendering the quiz.

## Project Setup
Ensure you have Node.js installed, then run the following commands to set up the project:

```bash
npm install
```

This installs all necessary dependencies.

## Running the App
To start the development server, run:

```bash
npm run dev
```

Visit `http://localhost:5173` (or the port provided in your terminal) to view the app.

## Structure
- `src/components/Quiz.vue`: This component fetches quiz questions from the backend and displays them using SurveyJS.
- `src/components/Result.vue`: Displays the quiz results, including the number of correct answers and compares the user's performance with previous attempts.
- `src/stores/quiz.js`: A Pinia store that manages the state of the quiz, including storing the questions and results.

## State Management
The Pinia store is used to manage the quiz state. It provides functions to set quiz questions and results, which are then used across the `Quiz` and `Result` components to display the quiz and results, respectively.

### Accessing State and Actions
To use the quiz store in components:

```javascript
import { useQuizStore } from '@/stores/quiz';

const quizStore = useQuizStore();
```

### Fetching and Storing Questions
Questions are fetched in `Quiz.vue` and stored in the Pinia store:

```javascript
quizStore.setQuestions(fetchedQuestions);
```

### Storing and Accessing Results

After completing the quiz, results are stored in the Pinia store:

```javascript
quizStore.setSurveyResults(surveyResults);
```

Results are accessed in `Result.vue` to display the user's performance.

## SurveyJS Integration
SurveyJS is used to render the quiz in `Quiz.vue`. The survey model is initialized with questions fetched from the backend and displayed to the user. Responses are collected and calculated at client side for evaluation.

## Deployment
To build the app for production, run:

```bash
npm run build
```

This creates a `dist` folder containing the production-ready files.

## Notes
- Ensure the backend server is running and accessible for the frontend to fetch questions and submit results.
- Adjust the API endpoints in `Quiz.vue` as necessary to match your backend setup.

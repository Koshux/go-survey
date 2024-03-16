# Quiz App Frontend

This frontend for the quiz app is built using Vue 3 with Vite. It features Pinia for state management, SurveyJS for rendering the quiz, and integrates `vue-i18n` for internationalization to enhance user experience.

## Project Setup

Ensure you have Node.js installed, then run the following commands to set up the project:

```bash
npm install
```

This command installs all necessary dependencies.

## Running the App

To start the development server, run:

```bash
npm run dev
```

Visit `http://localhost:5173` (or the port provided in your terminal) to view the app.

## Key Features

- **Multilingual Support**: Implemented using `vue-i18n`, allowing users to select their preferred language for the quiz.
- **Fuzzy Logic Performance Category**: Displays nuanced performance feedback from the API, categorizing user results into performance levels (e.g., "Good to Excellent") based on fuzzy logic evaluation.
- **SurveyJS Integration**: Utilizes SurveyJS for an interactive quiz experience, with support for internationalization to match the selected language in `vue-i18n`.

## Structure

- `src/components/Quiz.vue`: Fetches quiz questions from the backend and displays them using SurveyJS with internationalization support.
- `src/components/Result.vue`: Displays quiz results, including the fuzzy logic-derived performance category and the percentile rank.
- `src/stores/quiz.js`: Manages quiz state, including storing questions, user answers, and results.

## State Management

Pinia store is utilized to manage quiz state, offering methods to set quiz questions, user answers, and results, facilitating their use across components.

### Accessing State and Actions

In components, access the quiz store as follows:

```javascript
import { useQuizStore } from '@/stores/quiz';

const quizStore = useQuizStore();
```

### Fetching and Storing Questions & Results

Questions are fetched and stored in the Pinia store within `Quiz.vue`. Results, including user answers and the fuzzy logic performance category, are stored after quiz completion and accessed in `Result.vue`.

## SurveyJS and `vue-i18n` Integration

SurveyJS is configured to render the quiz in `Quiz.vue`, with model initialization based on the backend-supplied questions. `vue-i18n` integration ensures that quiz content is displayed according to the user's language preference.

## Deployment

For production build:

```bash
npm run build
```

Generates a `dist` folder with production-ready files.

## Nice to Have

- **Review Answers**: Enhance `Result.vue` to display the questions alongside the user's answers and the correct answers, providing insightful feedback on performance.
- **Dockerization**: Containerize the app for easier deployment and scalability. A Dockerfile and docker-compose.yml file could simplify setup on various environments.
- **Translations**: Translate the response of the fuzzified performance category to match the selected language in `vue-i18n`.
- **PWA Support**: Implement Progressive Web App features to enhance the app's offline capabilities and user experience.
- **Accessibility**: Ensure the app is accessible to all users, including those with disabilities, by adhering to WCAG guidelines.
- **Performance Optimization**: Optimize the app for performance, including lazy loading, code splitting, and caching.
- **Testing**: Implement unit tests for components and store methods to ensure the app's reliability and robustness.

## Notes

- Ensure the backend server is operational for the frontend to fetch questions and submit results.
- API endpoints in `Quiz.vue` might need adjustment to align with your backend configuration.

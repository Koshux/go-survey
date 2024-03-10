<template>
  <div>
    <h1>Quiz</h1>
    <div v-if="loading">Loading...</div>
    <div v-else id="surveyElement"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { SurveyModel } from 'survey-vue'
import 'survey-core/defaultV2.min.css'
import { useQuizStore } from '../stores/quiz'

const survey = ref(null)
const loading = ref(true)

watch(survey, (newValue) => {
  if (newValue) {
    newValue.render("surveyElement")
    newValue.onComplete.add((result) => {
      const quizStore = useQuizStore()
      quizStore.setResults(result.data)
    })
  }
}, { immediate: true })

const fetchQuestions = async () => {
  try {
    const response = await fetch('http://localhost:8080/questions')
    const questions = await response.json()
    const susrveyJson = convertQuestionsToSurveyJSFormat(questions)
    survey.value = new SurveyModel(susrveyJson)
  } catch (err) {
    console.error('Failed to fetch questions:', err)
  } finally {
    loading.value = false
  }
}

function convertQuestionsToSurveyJSFormat (questions) {
  return {
    title: 'Quiz',
    showProgressBar: 'bottom',
    pages: [
      {
        questions: questions.map((question, index) => ({
          type: 'radiogroup',
          name: question.id,
          // name: `question${index + 1}`,
          title: question.text,
          choices: question.options,
          correctAnswer: question.options[question.answer]
        })),
      },
    ],
  }
}

onMounted(fetchQuestions)
</script>

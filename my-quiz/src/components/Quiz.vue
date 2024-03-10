<template>
  <div>
    <div v-if="loading">Loading...</div>
    <div v-else-if="quizStore.survey">
      <SurveyComponent :model="quizStore.survey" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import 'survey-core/defaultV2.min.css'
import { useQuizStore } from '../stores/quiz'

// const survey = ref(null)
const loading = ref(true)
const quizStore = useQuizStore()

watch(quizStore.survey, (newValue) => {
  if (newValue) {
    console.log('Survey complete:', newValue)
    newValue.onComplete.add((result) => {
      console.log('Survey onComplete result:', result)
      console.log('Survey onComplete:', result.data)
      quizStore.setResults(result.data)
    })
  }
}, { immediate: true })

function convertQuestionsToSurveyJSFormat (questions) {
  return {
    title: 'General Knwoledge Quiz',
    showProgressBar: 'bottom',
    pages: [{
      elements: [{
        type: 'html',
        html: 'You are about to start a quiz on American history. <br>You will have 10 seconds for every question.<br>Enter your name below and click <b>Start Quiz</b> to begin.'
      }, {
        type: 'text',
        name: 'username',
        titleLocation: 'hidden',
        isRequired: true
      }],
    },
      ...questions.map((question, index) => ({
        elements: [{
          type: 'radiogroup',
          name: question.id,
          title: question.text,
          choices: question.options,
          correctAnswer: question.options[question.answer],
          requiredText: '!'
        }]
      }))
    ],
  }
}

onMounted(async () => {
  try {
    const response = await fetch('http://localhost:8080/questions')
    const questions = await response.json()
    quizStore.setQuestions(questions)

    const surveyJson = convertQuestionsToSurveyJSFormat(questions)
    quizStore.setSurvey(surveyJson)
  } catch (err) {
    console.error('Failed to fetch questions:', err)
  } finally {
    loading.value = false
  }
})
</script>

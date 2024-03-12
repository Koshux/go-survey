<template>
  <div>
    <div v-if="loading">Loading...</div>
    <div v-else>
      <SurveyComponent :model="survey" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import 'survey-core/defaultV2.min.css'
import { Model } from 'survey-core'
import { ContrastDarkPanelless } from 'survey-core/themes/contrast-dark-panelless'
import { useQuizStore } from '@/stores/quiz'

const survey = ref(null)
const loading = ref(true)
const quizStore = useQuizStore()

watch(survey, (newValue) => {
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
    title: 'General Knowledge Quiz',
    showProgressBar: 'bottom',
    pages: [
      ...questions.map((question, index) => ({
        elements: [{
          type: 'radiogroup',
          name: question.id,
          title: question.text,
          choices: question.options,
          correctAnswer: question.options[question.answer],
          isRequired: true
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
    survey.value = new Model(surveyJson)
    survey.value.startTimerOnFirstPage = false
    survey.value.showTimerPanel = 'top'
    survey.value.maxTimeToFinishPage = 10
    survey.value.applyTheme(ContrastDarkPanelless)
  } catch (err) {
    console.error('Failed to fetch questions:', err)
  } finally {
    loading.value = false
  }
})
</script>

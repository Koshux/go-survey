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
import 'survey-core/i18n/portuguese'
import 'survey-core/i18n/german'
import 'survey-core/i18n/swedish'
import { useQuizStore } from '@/stores/quiz'

const survey = ref(null)
const loading = ref(true)
const quizStore = useQuizStore()

watch(survey, (newValue) => {
  if (newValue) {
    newValue.onComplete.add((result) => {
      quizStore.setResults(result.data)
    })
  }
}, { immediate: true })

function convertQuestionsToSurveyJSFormat (questions) {
  return {
    title: 'General Knowledge Quiz',
    showProgressBar: 'bottom',
    showTimerPanel: "top",
    maxTimeToFinishPage: 10,
    firstPageIsStarted: true,
    startSurveyText: "Start Quiz",
    pages: [
      {
        elements: [{
          type: 'html',
          html: 'You are about to start a quiz on General Knowledge. <br>You will have 10 seconds for every question.<br>Enter your name below and click <b>Start Quiz</b> to begin.'
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
        }]
      }))
    ]
  }
}

// const setLocale = (locale) => {
//   survey.value.locale = locale
// }

// const showLo

onMounted(async () => {
  try {
    // TODO: create composable to send results to server
    const response = await fetch('http://localhost:8080/api/v1/quiz/questions')
    const questions = await response.json()
    quizStore.setQuestions(questions)

    const surveyJson = convertQuestionsToSurveyJSFormat(questions)
    survey.value = new Model(surveyJson)
    survey.value.startTimerOnFirstPage = false
    survey.value.showTimerPanel = 'top'
    survey.value.maxTimeToFinishPage = 10
    survey.value.applyTheme(ContrastDarkPanelless)
    survey.value.locale = 'pt'
  } catch (err) {
    console.error('Failed to fetch questions:', err)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div>
    <div v-if="isLoading">Loading...</div>
    <div v-else-if="error">{{ error }}</div>
    <div v-else>
      <template v-if="survey != null">
        <SurveyComponent :model="survey" />
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import 'survey-core/defaultV2.min.css'
import { Model } from 'survey-core'
import { ContrastDarkPanelless } from 'survey-core/themes/contrast-dark-panelless'
import 'survey-core/i18n/portuguese'
import 'survey-core/i18n/german'
import 'survey-core/i18n/swedish'

import { useQuizStore } from '@/stores/quiz'
import { useFetch } from '@/composables/useFetch'

const survey = ref(null)
const quizStore = useQuizStore()

const { data, error, execute, isLoading }= useFetch()
const { t } = useI18n()

watch(survey, (newValue) => {
  if (newValue) {
    newValue.onComplete.add((result) => {
      const res = {}
      res.username = result.data.username
      res.answers = {}

      // Convert the result data to the API submit answers format.
      Object.keys(result.data).map(key => {
        if (isNaN(parseInt(key))) {
          res.username = result.data[key]
        } else {
          res.answers[key] = result.data[key]
        }
      })

      quizStore.setSurveyResults(res)
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
    startSurveyText: t('start'),
    pages: [
      {
        elements: [{
          type: 'html',
          html: t('quiz.starter_page')
        }, {
          type: 'text',
          name: 'username',
          titleLocation: 'hidden',
          isRequired: true
        }],
      },
      ...questions.map((question, index) => {
        return ({
          elements: [{
            type: 'radiogroup',
            name: question.id,
            title: question.text,
            choices: question.options,
            correctAnswer: question.options[question.answer],
          }]
        })
      })
    ]
  }
}

onMounted(async () => {
  await execute(
    'http://localhost:8080/api/v1/quiz/questions',
    { method: 'GET', headers: { 'Content-Type': 'application/json' } }
  )
  quizStore.setQuestions(data.value)

  const surveyJson = convertQuestionsToSurveyJSFormat(data.value)
  survey.value = new Model(surveyJson)
  survey.value.startTimerOnFirstPage = false
  survey.value.showTimerPanel = 'top'
  survey.value.maxTimeToFinishPage = 10
  survey.value.applyTheme(ContrastDarkPanelless)
  quizStore.setSurveyModel(survey.value)
})
</script>

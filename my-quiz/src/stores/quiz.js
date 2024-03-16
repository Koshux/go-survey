import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'
import { useI18n } from 'vue-i18n'

export const useQuizStore = defineStore('quiz', () => {
  const questions = ref([])
  const quizResponse = ref({})
  const surveyResults = ref([])
  const survey = ref(null)
  const currentLanguage = ref('en')
  const { locale, t } = useI18n()

  const resultCount = computed(() => {
    return surveyResults.value != null ? apiScore : 0
  })

  const completed = computed(() => {
    return surveyResults.value != null && Object.keys(surveyResults.value).length > 0
  })

  const correctAnswers = computed(() => {
    if (quizResponse.value == null) return 0

    const { score } = quizResponse.value

    return score != null ? score : 0
  })

  const percentile = computed(() => {
    const { percentile } = quizResponse.value

    if (percentile == null) return t('result.score_invalid')
    if (percentile === 999) return t('result.score_insufficient')
    return t('result.score_ranked', {
      percentile: percentile != null ? percentile : 0
    })
  })

  const username = computed(() => {
    const { username } = surveyResults.value
    console.log('survey results', surveyResults.value, username)

    return username != null ? username : 'Anonymous'
  })

  watch(() => currentLanguage.value, (language) => {
    if (survey.value != null) {
      console.log('Language changed to', language)
      survey.value.locale = language
      locale.value = language
    }
  })

  function setSurveyModel (model) {
    survey.value = model
  }

  function setSurveyResults (data) {
    surveyResults.value = data
  }

  function setQuizResponse (data) {
    quizResponse.value = data
  }

  function setQuestions (data) {
    questions.value = data
  }

  function reset () {
    surveyResults.value = []
  }

  function setLanguage (language) {
    currentLanguage.value = language
  }

  return {
    completed,
    correctAnswers,
    currentLanguage,
    percentile,
    questions,
    reset,
    surveyResults,
    resultCount,
    survey,
    username,
    setQuestions,
    setQuizResponse,
    setSurveyResults,
    setLanguage,
    setSurveyModel,
  }
})

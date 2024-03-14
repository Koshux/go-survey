import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

export const useQuizStore = defineStore('quiz', () => {
  const questions = ref([])
  const quizResult = ref({})
  const results = ref([])
  const survey = ref(null)
  const currentLanguage = ref('en')

  const resultCount = computed(() => {
    return results.value != null ? Object.keys(results.value).length : 0
  })

  const completed = computed(() => {
    return results.value != null && Object.keys(results.value).length > 0
  })

  const correctAnswers = computed(() => {
    const { username, ...answers } = results.value

    return Object.keys(answers).reduce((acc, questionId) => {
      if (!questions.value.length) return acc

      const question = questions.value.find((q) => q.id === questionId)

      if (!question) return acc

      const correctAnswer = question.options[question.answer]
      const userAnswer = answers[questionId]

      return correctAnswer === userAnswer ? acc + 1 : acc
    }, 0)
  })

  watch(() => currentLanguage.value, (language) => {
    if (survey.value != null) {
      console.log('Language changed to', language)
      survey.value.locale = language
    }
  })

  function setSurveyModel (model) {
    survey.value = model
  }

  function setQuizResult (data) {
    quizResult.value = data
  }

  function setResults (data) {
    results.value = data
  }

  function setQuestions (data) {
    questions.value = data
  }

  function reset () {
    results.value = []
  }

  function setLanguage (language) {
    currentLanguage.value = language
  }

  return {
    completed,
    correctAnswers,
    currentLanguage,
    questions,
    reset,
    results,
    resultCount,
    survey,
    setQuestions,
    setQuizResult,
    setResults,
    setLanguage,
    setSurveyModel,
  }
})

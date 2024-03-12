import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const useQuizStore = defineStore('quiz', () => {
  const questions = ref([])
  const results = ref([])
  const survey = ref(null)

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

  function setResults (data) {
    results.value = data
  }

  function setQuestions (data) {
    questions.value = data
  }

  function reset () {
    results.value = []
  }

  return {
    completed,
    correctAnswers,
    questions,
    reset,
    results,
    resultCount,
    survey,
    setQuestions,
    setResults,
  }
})

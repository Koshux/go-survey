import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const useQuizStore = defineStore('quiz', () => {
  const results = ref(null)
  const survey = ref(null)

  const completed = computed(() => !!results.value)

  function setSurvey (data) {
    survey.value = data
  }

  function setResults (data) {
    results.value = data
  }

  return { completed, results, survey, setSurvey, setResults }
})

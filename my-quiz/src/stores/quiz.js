import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { Model } from 'survey-core'
import { ContrastDarkPanelless } from 'survey-core/themes/contrast-dark-panelless'

export const useQuizStore = defineStore('quiz', () => {
  const questions = ref([])
  const results = ref([])
  const survey = ref(null)

  const completed = computed(() => {
    return results.value != null && results.value.length > 0
  })

  function setSurvey (data) {
    survey.value = new Model(data)
    survey.value.startTimerOnFirstPage = false
    survey.value.showTimerPanel = 'top'
    survey.value.maxTimeToFinishPage = 10
    survey.value.applyTheme(ContrastDarkPanelless)
  }

  function setSurveySetting(settingName, data) {
    survey.value[settingName] = data
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

  return {
    completed,
    questions,
    reset,
    results,
    survey,
    setQuestions,
    setResults,
    setSurveySetting,
    setSurvey
  }
})

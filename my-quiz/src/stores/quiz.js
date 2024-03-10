import { defineStore } from 'pinia'

export const useQuizStore = defineStore('quiz', {
  state: () => ({
    results: null
  }),
  actions: {
    setResults(results) {
      this.results = results
    }
  }
})

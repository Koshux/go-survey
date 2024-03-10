<template>
  <div>
    <h1>Result</h1>
    <p v-if="loading">Loading...</p>
    <div v-else>
      <p>You got {{ correctAnswers }} out of {{ quizStore.results.length }} questions right!</p>
      <button @click="reset">Try again</button>
    </div>
  </div>
</template>

<script setup>
import { computed, watch, ref } from 'vue';
import { useQuizStore } from '../stores/quiz'

const quizStore = useQuizStore()

const loading = ref(true)
const correctAnswers = ref(0)

const results = computed(() => quizStore.results || [])

// I need to display the correct number of answers once the results are available
watch(results, (newResults, oldResults) => {
  if (newResults !== oldResults) {
    debugger
    loading.value = false
    console.log('newResults:', newResults.value)
    correctAnswers.value = newResults.value.reduce((acc, answer) => {
      const question = quizStore.questions.value.find((q) => q.id === answer.name)
      return question.options[question.answer] === answer.value ? acc + 1 : acc
    }, 0)
  }
}, { deep: true, immediate: true })
</script>

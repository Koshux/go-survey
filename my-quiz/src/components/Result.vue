<template>
  <div>
    <h1>Thank you, {{ username }}!</h1>
    <p v-if="isLoading">Loading...</p>
    <p v-else-if="error">{{ error }}</p>
    <div v-else>
      <p>You got {{ correctAnswers }} out of {{ questions.length }} questions right!</p>
      <p>{{ percentile }}</p>
      <button @click="reset">Try again</button>
    </div>
  </div>
</template>

<script setup>
import { watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useQuizStore } from '@/stores/quiz'
import { useFetch } from '../composables/useFetch';

const quizStore = useQuizStore()
const { data: answers, error, execute, isLoading }= useFetch()
const {
  correctAnswers,
  percentile,
  questions,
  surveyResults,
  username,
} = storeToRefs(quizStore)

function reset () {
  quizStore.reset()
}

watch(surveyResults, async (newResults, oldResults) => {
  if (newResults !== oldResults) {
    await execute(
      'http://localhost:8080/api/v1/quiz/answers',
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(surveyResults.value)
      }
    )

    quizStore.setQuizResponse(answers.value)
  }
}, { deep: true, immediate: true })
</script>

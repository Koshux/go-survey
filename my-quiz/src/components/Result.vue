<template>
  <div>
    <h1>Result</h1>
    <p v-if="isLoading">Loading...</p>
    <div v-else>
      <p>You got {{ correctAnswers }} out of {{ resultCount }} questions right!</p>
      <button @click="reset">Try again</button>
    </div>
  </div>
</template>

<script setup>
import { watch, ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useQuizStore } from '@/stores/quiz'
import { useFetch } from '../composables/useFetch';

const quizStore = useQuizStore()
const { data: answers, error, execute, isLoading }= useFetch()
const { correctAnswers, results, resultCount } = storeToRefs(quizStore)

function reset () {
  quizStore.reset()
}

watch(correctAnswers, async (newResults, oldResults) => {
  if (newResults !== oldResults) {
    // I need to map the json object of quizStore.results to the format expected by the backend.
    // The object is in the format { "Username": "asda", "Answers": { "1": "Paris", "2": "Mars", "3": "Pacific Ocean", "4": "1912", "5": "William Shakespeare" } }.
    // It keeps going until 10 questions. The quizStore.results is the format like the object above for the property "Answers".
    const quizResult = await execute(
      'http://localhost:8080/api/v1/quiz/answers',
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(results.value)
      }
    )
    quizStore.setQuizResult(quizResult)
    console.log('quizResult:', quizResult)
  }
}, { deep: true, immediate: true })
</script>

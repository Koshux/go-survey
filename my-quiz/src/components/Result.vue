<template>
  <div>
    <h1>Result</h1>
    <p v-if="loading">Loading...</p>
    <div v-else>
      <p>You got {{ correctAnswers }} out of {{ questions.length }} questions right!</p>
      <button @click="reset">Try again</button>
    </div>
  </div>
</template>

<script setup>
import { useQuizStore } from '../stores/quiz'

const quizStore = useQuizStore()
const results = quizStore.results

// newValue.onComplete.add((result) => {
//   const correctAnswers = result.data.reduce((acc, answer) => {
//     const question = questions.value.find((q) => q.id === answer.name)
//     return question.options[question.answer] === answer.value ? acc + 1 : acc
//   }, 0)
//   console.log('correctAnswers:', correctAnswers)
// })
</script>

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

const correctAnswers = computed(() => {
  const { username, ...answers } = quizStore.results

  return Object.keys(answers).reduce((acc, questionId) => {
    const question = quizStore.questions.find((q) => q.id === questionId)
    if (!question) return acc

    const correctAnswer = question.options[question.answer]
    const userAnswer = answers[questionId]

    return correctAnswer === userAnswer ? acc + 1 : acc
  }, 0)
})

function reset () {
  quizStore.reset()
}

watch(correctAnswers, (newResults, oldResults) => {
  if (newResults !== oldResults) {
    // debugger
    loading.value = false
    console.log('newResults:', JSON.stringify(JSON.parse(newResults)))

    quizStore.setSurveySetting('completedHtml', `
      <h4>You got <b>${correctAnswers}</b> out of <b>${quizStore.results.length}</b> correct answers.</h4>
    `)

    quizStore.setSurveySetting('completedHtmlOnCondition', [{
        expression: `${correctAnswers} == 0`,
        html: "<h4>Unfortunately, none of your answers are correct. Please try again.</h4>"
    }, {
        expression: `${correctAnswers} == ${quizStore.results.length}`,
        html: "<h4>Congratulations! You answered all the questions correctly!</h4>"
    }])
  }
}, { deep: true, immediate: true })
</script>

<template>
  <div>
    <h1>Result</h1>
    <p v-if="loading">Loading...</p>
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

const quizStore = useQuizStore()
const { correctAnswers, resultCount } = storeToRefs(quizStore)
const loading = ref(true)

function reset () {
  quizStore.reset()
}

watch(correctAnswers, async (newResults, oldResults) => {
  if (newResults !== oldResults) {
    loading.value = false
    console.log('newResults:', JSON.stringify(JSON.parse(newResults)))
    // TODO: create composable to send results to server
    const quizResult = await fetch('http://localhost:8080/api/v1/quiz/answers', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(JSON.parse(newResults))
    })
    quizStore.setQuizResult(await quizResult.json())
    console.log('quizResult:', quizResult)
    // quizStore.setSurveySetting('completedHtml', `
    //   <h4>You got <b>${correctAnswers}</b> out of <b>${quizStore.results.length}</b> correct answers.</h4>
    // `)

    // quizStore.setSurveySetting('completedHtmlOnCondition', [{
    //     expression: `${correctAnswers} == 0`,
    //     html: "<h4>Unfortunately, none of your answers are correct. Please try again.</h4>"
    // }, {
    //     expression: `${correctAnswers} == ${quizStore.results.length}`,
    //     html: "<h4>Congratulations! You answered all the questions correctly!</h4>"
    // }])
  }
}, { deep: true, immediate: true })
</script>

<template>
  <div>
    <h1>{{ t('result.title', { username }) }}</h1>
    <p v-if="isLoading">Loading...</p>
    <p v-else-if="error">{{ error }}</p>
    <div v-else>
      <p>{{ score }}</p>
      <p>{{ percentile }}</p>
      <button @click="reset">{{ t('result.reset') }}</button>
    </div>
  </div>
</template>

<script setup>
import { watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useQuizStore } from '@/stores/quiz'
import { useFetch } from '../composables/useFetch'
import { useI18n } from 'vue-i18n'

const quizStore = useQuizStore()
const { t } = useI18n()
const { data: answers, error, execute, isLoading }= useFetch()
const {
  percentile,
  score,
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

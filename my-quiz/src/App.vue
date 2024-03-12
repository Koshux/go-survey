<script setup>
import { ref } from 'vue'
import Quiz from '@/components/Quiz.vue'
import Result from '@/components/Result.vue'
import { useQuizStore } from '@/stores/quiz'

const quizStore = useQuizStore()
</script>

<template>
  <div>
    <Quiz v-if="!quizStore.completed" />
    <Result v-else />
  </div>
</template>

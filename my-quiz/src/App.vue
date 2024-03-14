<script setup>
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import Quiz from '@/components/Quiz.vue'
import Result from '@/components/Result.vue'
import { useQuizStore } from '@/stores/quiz'

const quizStore = useQuizStore()
const { currentLanguage } = storeToRefs(quizStore)
const selectedLanguage = ref(currentLanguage)
const { completed } = storeToRefs(quizStore)

function updateLanguage () {
  quizStore.setLanguage(selectedLanguage.value)
}
</script>

<template>
  <div>
    <div v-if="!completed" class="mb-2">
      <select v-model="selectedLanguage" @change="updateLanguage">
        <option value="en">English</option>
        <option value="de">German</option>
        <option value="pt">Portuguese</option>
        <option value="sv">Swedish</option>
      </select>
    </div>

    <Quiz v-if="!completed" />
    <Result v-else />
  </div>
</template>


<style scoped>
.mb-2 {
  margin-bottom: 1rem;
}

</style>

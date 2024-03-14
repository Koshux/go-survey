<script setup>
import { ref } from 'vue'
import Quiz from '@/components/Quiz.vue'
import Result from '@/components/Result.vue'
import { useQuizStore } from '@/stores/quiz'
import { storeToRefs } from 'pinia';

const quizStore = useQuizStore()
const { currentLanguage } = storeToRefs(quizStore)
const selectedLanguage = ref(currentLanguage)

function updateLanguage () {
  quizStore.setLanguage(selectedLanguage.value)
}
</script>

<template>
  <div>
    <div class="mb-2">
      <select v-model="selectedLanguage" @change="updateLanguage">
        <option value="en">English</option>
        <option value="de">German</option>
        <option value="pt">Portuguese</option>
        <option value="sv">Swedish</option>
      </select>
    </div>

    <Quiz v-if="!quizStore.completed" />
    <Result v-else />
  </div>
</template>


<style scoped>
.mb-2 {
  margin-bottom: 1rem;
}

</style>

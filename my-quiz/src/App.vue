<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import Quiz from '@/components/Quiz.vue'
import Result from '@/components/Result.vue'
import { useQuizStore } from '@/stores/quiz'

const quizStore = useQuizStore()
const { currentLanguage } = storeToRefs(quizStore)
const { completed } = storeToRefs(quizStore)

const selectedLanguage = ref(currentLanguage)

const { t } = useI18n()

function updateLanguage () {
  quizStore.setLanguage(selectedLanguage.value)
}
</script>

<template>
  <div>
    <div class="mb-2">
      <select v-model="selectedLanguage" @change="updateLanguage">
        <option value="en">{{ t('en') }}</option>
        <option value="de">{{ t('de') }}</option>
        <option value="pt">{{ t('pt') }}</option>
        <option value="sv">{{ t('sv') }}</option>
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

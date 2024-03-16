import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { surveyPlugin } from 'survey-vue3-ui'
import './style.css'
import App from './App.vue'
import i18n from './i18n'

createApp(App)
  .use(i18n)
  .use(createPinia())
  .use(surveyPlugin)
  .mount('#app')

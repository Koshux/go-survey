import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { surveyPlugin } from 'survey-vue3-ui'
import './style.css'
import App from './App.vue'

createApp(App)
  .use(createPinia())
  .use(surveyPlugin)
  .mount('#app')

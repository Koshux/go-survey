import { createI18n } from 'vue-i18n'

const messages = {
  en: {
    welcome: 'Welcome to my quiz app!',
    start: 'Start Quiz',
    anonymous: 'Anonymous',
    language: 'Language',
    en: 'English',
    de: 'German',
    pt: 'Portuguese',
    sv: 'Swedish',
    quiz: {
      welcome: 'Welcome to the quiz!',
      start: 'Start Quiz',
    },
    result: {
      reset: 'Try again',
      score_invalid: 'You didn\'t score any points.',
      score_insufficient: 'There are insufficient data to calculate your score, please check back later.',
      score_ranked: 'You scored better than {percentile}% of the quiz takers.'
    },
  },
  de: {

  },
  pt: {

  },
  sv: {

  }
}

const i18n = createI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  messages,
})

export default i18n

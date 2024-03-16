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
    welcome: 'Willkommen zu meiner Quiz-App!',
    start: 'Quiz starten',
    anonymous: 'Anonym',
    language: 'Sprache',
    en: 'Englisch',
    de: 'Deutsch',
    pt: 'Portugiesisch',
    sv: 'Schwedisch',
    quiz: {
      welcome: 'Willkommen zum Quiz!',
      start: 'Quiz starten',
    },
    result: {
      reset: 'Nochmal versuchen',
      score_invalid: 'Du hast keine Punkte erzielt.',
      score_insufficient: 'Es gibt nicht genügend Daten, um deinen Score zu berechnen. Bitte versuche es später erneut.',
      score_ranked: 'Du hast besser als {percentile}% der Quizteilnehmer abgeschnitten.'
    },
  },
  pt: {
    welcome: 'Bem-vindo ao meu aplicativo de quiz!',
    start: 'Começar Quiz',
    anonymous: 'Anônimo',
    language: 'Idioma',
    en: 'Inglês',
    de: 'Alemão',
    pt: 'Português',
    sv: 'Sueco',
    quiz: {
      welcome: 'Bem-vindo ao quiz!',
      start: 'Começar Quiz',
    },
    result: {
      reset: 'Tentar novamente',
      score_invalid: 'Você não marcou nenhum ponto.',
      score_insufficient: 'Não há dados suficientes para calcular sua pontuação, por favor, verifique mais tarde.',
      score_ranked: 'Você marcou melhor do que {percentile}% dos participantes do quiz.'
    },
  },
  sv: {
    welcome: 'Välkommen till min quiz-app!',
    start: 'Starta Quiz',
    anonymous: 'Anonym',
    language: 'Språk',
    en: 'Engelska',
    de: 'Tyska',
    pt: 'Portugisiska',
    sv: 'Svenska',
    quiz: {
      welcome: 'Välkommen till quizet!',
      start: 'Starta Quiz',
    },
    result: {
      reset: 'Försök igen',
      score_invalid: 'Du har inte fått några poäng.',
      score_insufficient: 'Det finns otillräckliga data för att beräkna din poäng, vänligen kontrollera senare.',
      score_ranked: 'Du fick bättre än {percentile}% av quizdeltagarna.'
    },
  }
}

const i18n = createI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  messages,
})

export default i18n

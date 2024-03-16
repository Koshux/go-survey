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
      starter_page: 'You are about to start a quiz on General Knowledge. <br>You will have 10 seconds for every question.<br>Enter your name below and click <b>Start Quiz</b> to begin.',
    },
    result: {
      title: 'Thank you, {username}!',
      reset: 'Try again',
      score: 'You got {correct} out of {total} questions right!',
      percentile_invalid: 'You didn\'t score any points.',
      percentile_insufficient: 'There are insufficient data to calculate your score, please check back later.',
      percentile_ranked: 'You scored better than {percentile}% of the quiz takers.'
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
      starter_page: 'Du bist dabei, ein Quiz über Allgemeinwissen zu starten. <br>Du hast 10 Sekunden für jede Frage.<br>Gib deinen Namen unten ein und klicke auf <b>Quiz starten</b>, um zu beginnen.',
    },
    result: {
      title: 'Danke, {username}!',
      reset: 'Nochmal versuchen',
      score: 'Du hast {correct} von {total} Fragen richtig beantwortet!',
      percentile_invalid: 'Du hast keine Punkte erzielt.',
      percentile_insufficient: 'Es gibt nicht genügend Daten, um deinen Score zu berechnen. Bitte versuche es später erneut.',
      percentile_ranked: 'Du hast besser als {percentile}% der Quizteilnehmer abgeschnitten.'
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
      starter_page: 'Você está prestes a começar um quiz sobre conhecimentos gerais. <br>Você terá 10 segundos para cada pergunta.<br>Insira seu nome abaixo e clique em <b>Começar Quiz</b> para começar.',
    },
    result: {
      title: 'Obrigado, {username}!',
      reset: 'Tentar novamente',
      score: 'Você acertou {correct} de {total} perguntas!',
      percentile_invalid: 'Você não marcou nenhum ponto.',
      percentile_insufficient: 'Não há dados suficientes para calcular sua pontuação, por favor, verifique mais tarde.',
      percentile_ranked: 'Você marcou melhor do que {percentile}% dos participantes do quiz.'
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
      starter_page: 'Du är på väg att starta ett quiz om Allmänbildning. <br>Du har 10 sekunder på dig för varje fråga.<br>Ange ditt namn nedan och klicka på <b>Starta Quiz</b> för att börja.',
    },
    result: {
      title: 'Tack, {username}!',
      reset: 'Försök igen',
      score: 'Du fick {correct} av {total} frågor rätt!',
      percentile_invalid: 'Du har inte fått några poäng.',
      percentile_insufficient: 'Det finns otillräckliga data för att beräkna din poäng, vänligen kontrollera senare.',
      percentile_ranked: 'Du fick bättre än {percentile}% av quizdeltagarna.'
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

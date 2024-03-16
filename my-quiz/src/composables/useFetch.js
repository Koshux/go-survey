import { ref } from 'vue'

export function useFetch () {
  const data = ref(null)
  const error = ref(null)
  const loading = ref(true)

  async function execute (url, options = {}) {
    loading.value = true
    error.value = null
    data.value = null

    try {
      console.log('Fetching data from', url);
      const response = await fetch(url, options)

      if (!response.ok) {
        throw new Error('An error occurred while fetching the data.')
      }

      const jsonData = await response.json()
      data.value = jsonData
    } catch (err) {
      error.value = err
    } finally {
      loading.value = false
    }
  }

  return { data, error, execute, loading }
}

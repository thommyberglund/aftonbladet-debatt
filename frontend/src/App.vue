<template>
  <div class="min-h-screen bg-gray-100 py-8 px-4 font-sans">
    <div class="max-w-4xl mx-auto">
      
      <header class="mb-10 text-center">
        <h1 class="text-4xl font-extrabold text-gray-900 mb-2">Aftonbladet Debatt</h1>
        <p class="text-gray-600">Senaste inläggen hämtade via REST API</p>
        <button 
          @click="fetchDebates" 
          :disabled="loading"
          class="mt-6 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-6 rounded-full transition duration-300 ease-in-out shadow-md disabled:opacity-50"
        >
          {{ loading ? 'Hämtar...' : 'Uppdatera flöde' }}
        </button>
      </header>

      <div v-if="loading" class="flex justify-center my-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <div v-else-if="error" class="bg-red-100 border-l-4 border-red-500 text-red-700 p-4 mb-6 shadow-sm">
        <p class="font-bold">Ett fel uppstod</p>
        <p>{{ error }}</p>
      </div>

      <div v-else class="grid gap-6">
        <div 
          v-for="(article, index) in debates" 
          :key="index" 
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
        >
          <div class="p-6">
            <div class="flex justify-between items-start mb-2">
              <h2 class="text-xl font-bold text-gray-800 leading-tight flex-1">
                {{ article.title }}
              </h2>
              <span class="text-xs text-gray-400 ml-4 whitespace-nowrap italic">
                {{ formatDate(article.published) }}
              </span>
            </div>
            
            <p class="text-gray-600 mb-4 leading-relaxed" v-html="article.summary"></p>
            
            <a 
              :href="article.link" 
              target="_blank" 
              class="inline-block text-blue-600 font-medium hover:text-blue-800 hover:underline transition-colors"
            >
              Läs hela debattartikeln →
            </a>
          </div>
        </div>
      </div>

      <div v-if="!loading && debates.length === 0 && !error" class="text-center text-gray-500 py-12">
        Inga artiklar hittades.
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const debates = ref([])
const loading = ref(false)
const error = ref(null)

const fetchDebates = async () => {
  loading.value = true
  error.value = null
  try {
    // Justera URL:en om din backend körs på en annan port
    const response = await fetch('http://127.0.0.1:8000/debatt')
    if (!response.ok) throw new Error('Kunde inte nå backend-tjänsten')
    
    const data = await response.json()
    debates.value = data.articles
  } catch (err) {
    error.value = err.message
    console.error("Fel vid hämtning:", err)
  } finally {
    loading.value = false
  }
}

const formatDate = (dateStr) => {
  if (dateStr === 'Information saknas') return dateStr
  // Enkel formatering för att snygga till datumet
  try {
    const date = new Date(dateStr)
    return date.toLocaleDateString('sv-SE', { day: 'numeric', month: 'short' })
  } catch {
    return dateStr
  }
}

onMounted(fetchDebates)
</script>

<style>
/* Om du inte använder Tailwind CLI kan du lägga till CDN i index.html */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;600;800&display=swap');
body { font-family: 'Inter', sans-serif; }
</style>

<template>
  <div class="container">
    <header class="header">
      <h1>Debatt-kollen</h1>
      <p>Senaste debatterna från Aftonbladet och Expressen</p>
      <button @click="fetchAll" :disabled="loading" class="refresh-btn">
        {{ loading ? 'Uppdaterar...' : 'Uppdatera båda flöden' }}
      </button>
    </header>

    <div class="grid-layout">
      <section class="column">
        <h2 class="source-title ab">Aftonbladet Debatt</h2>
        <div v-if="errorAB" class="error">{{ errorAB }}</div>
        <div v-for="a in abArticles" :key="a.link" class="card">
          <h3>{{ a.title }}</h3>
          <span class="date">{{ formatDate(a.published) }}</span>
          <p v-html="a.summary"></p>
          <a :href="a.link" target="_blank">Läs på AB →</a>
        </div>
      </section>

      <section class="column">
        <h2 class="source-title exp">Expressen Debatt</h2>
        <div v-if="errorExp" class="error">{{ errorExp }}</div>
        <div v-for="a in expArticles" :key="a.link" class="card">
          <h3>{{ a.title }}</h3>
          <span class="date">{{ formatDate(a.published) }}</span>
          <p v-html="a.summary"></p>
          <a :href="a.link" target="_blank">Läs på Expressen →</a>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const abArticles = ref([])
const expArticles = ref([])
const loading = ref(false)
const errorAB = ref(null)
const errorExp = ref(null)

const fetchAll = async () => {
  loading.value = true
  // Hämta båda parallellt
  await Promise.all([fetchAB(), fetchExp()])
  loading.value = false
}

const fetchAB = async () => {
  try {
    const res = await fetch('/api/debatt')
    const data = await res.json()
    abArticles.value = data.articles
  } catch (e) { errorAB.value = "Kunde inte ladda Aftonbladet" }
}

const fetchExp = async () => {
  try {
    const res = await fetch('/api/expressen')
    const data = await res.json()
    expArticles.value = data.articles
  } catch (e) { errorExp.value = "Kunde inte ladda Expressen" }
}

const formatDate = (d) => new Date(d).toLocaleDateString('sv-SE')

onMounted(fetchAll)
</script>

<style scoped>
.container { max-width: 1200px; margin: 0 auto; padding: 20px; font-family: sans-serif; }
.header { text-align: center; margin-bottom: 40px; }
.refresh-btn { padding: 10px 25px; border-radius: 20px; border: none; bg-color: #333; color: white; cursor: pointer; }

.grid-layout { 
  display: grid; 
  grid-template-columns: 1fr 1fr; 
  gap: 30px; 
}

@media (max-width: 768px) {
  .grid-layout { grid-template-columns: 1fr; }
}

.column { background: #f9f9f9; padding: 20px; border-radius: 12px; }
.source-title { padding-bottom: 10px; border-bottom: 4px solid; margin-bottom: 20px; }
.ab { border-color: #e01010; } /* Röd för AB */
.exp { border-color: #005696; } /* Blå för Expressen */

.card { background: white; padding: 15px; border-radius: 8px; margin-bottom: 15px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
.card h3 { margin: 0 0 5px 0; font-size: 1.1rem; }
.date { font-size: 0.8rem; color: #888; }
.card a { display: inline-block; margin-top: 10px; text-decoration: none; font-weight: bold; color: #007bff; }
.error { color: red; font-weight: bold; }
</style>
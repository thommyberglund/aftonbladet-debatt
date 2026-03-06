<template>
  <div class="app-wrapper">
    <header class="main-header">
      <div class="header-content">
        <h1>Debattkollen</h1>
        <p class="subtitle">Senaste opinionen från Sveriges största redaktioner</p>
        <button @click="fetchAll" :disabled="loading" class="btn-refresh">
          {{ loading ? 'Uppdaterar...' : 'Uppdatera alla flöden' }}
        </button>
      </div>
    </header>

    <main class="grid-container">
      <section class="feed-column">
        <div class="column-header">
          <div class="brand-badge ab">Aftonbladet</div>
        </div>
        <div v-if="errorAB" class="status-msg error">{{ errorAB }}</div>
        <div class="articles-list">
          <article v-for="a in abArticles" :key="a.link" class="article-card">
            <span class="date">{{ formatDate(a.published) }}</span>
            <h3 class="title">{{ a.title }}</h3>
            <p class="summary" v-html="truncate(a.summary)"></p>
            <a :target="_blank" :href="a.link" class="read-more">Läs på AB</a>
          </article>
        </div>
      </section>

      <section class="feed-column">
        <div class="column-header">
          <div class="brand-badge exp">Expressen</div>
        </div>
        <div v-if="errorExp" class="status-msg error">{{ errorExp }}</div>
        <div class="articles-list">
          <article v-for="a in expArticles" :key="a.link" class="article-card">
            <span class="date">{{ formatDate(a.published) }}</span>
            <h3 class="title">{{ a.title }}</h3>
            <p class="summary" v-html="truncate(a.summary)"></p>
            <a :target="_blank" :href="a.link" class="read-more">Läs på Expressen</a>
          </article>
        </div>
      </section>

      <section class="feed-column">
        <div class="column-header">
          <div class="brand-badge dn">DN Debatt</div>
        </div>
        <div v-if="errorDN" class="status-msg error">{{ errorDN }}</div>
        <div class="articles-list">
          <article v-for="a in dnArticles" :key="a.link" class="article-card">
            <span class="date">{{ formatDate(a.published) }}</span>
            <h3 class="title">{{ a.title }}</h3>
            <p class="summary" v-html="truncate(a.summary)"></p>
            <a :target="_blank" :href="a.link" class="read-more">Läs på DN</a>
          </article>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const abArticles = ref([])
const expArticles = ref([])
const dnArticles = ref([])
const loading = ref(false)
const errorAB = ref(null), errorExp = ref(null), errorDN = ref(null)

const fetchAll = async () => {
  loading.value = true
  await Promise.all([fetchAB(), fetchExp(), fetchDN()])
  loading.value = false
}

const fetchAB = async () => {
  try {
    const res = await fetch('/api/debatt')
    const data = await res.json()
    abArticles.value = data.articles
  } catch (e) { errorAB.value = "Fel vid laddning av AB" }
}

const fetchExp = async () => {
  try {
    const res = await fetch('/api/expressen')
    const data = await res.json()
    expArticles.value = data.articles
  } catch (e) { errorExp.value = "Fel vid laddning av Expressen" }
}

const fetchDN = async () => {
  try {
    const res = await fetch('/api/dn')
    const data = await res.json()
    dnArticles.value = data.articles
  } catch (e) { errorDN.value = "Fel vid laddning av DN" }
}

const formatDate = (d) => d && d !== 'Saknas' ? new Date(d).toLocaleDateString('sv-SE', { day: 'numeric', month: 'short' }) : ''
const truncate = (text) => text ? (text.length > 140 ? text.substring(0, 140) + '...' : text) : ''

onMounted(fetchAll)
</script>

<style scoped>
.app-wrapper { background: #f8fafc; min-height: 100vh; font-family: sans-serif; padding-bottom: 40px; }
.main-header { background: white; padding: 30px; border-bottom: 1px solid #e2e8f0; text-align: center; margin-bottom: 30px; }
.grid-container { 
  max-width: 1400px; 
  margin: 0 auto; 
  display: grid; 
  grid-template-columns: repeat(3, 1fr); 
  gap: 20px; 
  padding: 0 20px; 
}

@media (max-width: 1024px) { grid-container { grid-template-columns: 1fr; } }

.brand-badge { font-weight: bold; padding: 4px 12px; border-radius: 4px; font-size: 0.8rem; }
.ab { background: #fee2e2; color: #b91c1c; }
.exp { background: #dbeafe; color: #1e40af; }
.dn { background: #f3f4f6; color: #111827; border: 1px solid #d1d5db; }

.article-card { background: white; border: 1px solid #e2e8f0; border-radius: 8px; padding: 16px; margin-bottom: 16px; }
.title { font-size: 1.1rem; margin: 8px 0; line-height: 1.3; }
.date { font-size: 0.7rem; color: #94a3b8; font-weight: bold; }
.summary { font-size: 0.9rem; color: #475569; margin-bottom: 12px; line-height: 1.5; }
.read-more { font-size: 0.85rem; color: #2563eb; text-decoration: none; font-weight: 600; }
.btn-refresh { background: #111827; color: white; border: none; padding: 10px 20px; border-radius: 20px; cursor: pointer; }
</style>
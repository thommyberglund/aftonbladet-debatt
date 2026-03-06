<template>
  <div class="app-wrapper">
    <header class="main-header">
      <div class="header-content">
        <div class="logo">
          <span class="logo-icon">🗞️</span>
          <h1>Debattkollen</h1>
        </div>
        <p class="subtitle">Senaste opinionen från Sveriges största redaktioner</p>
        <button @click="fetchAll" :disabled="loading" class="btn-refresh">
          <span v-if="loading" class="spinner"></span>
          {{ loading ? 'Uppdaterar...' : 'Uppdatera flöden' }}
        </button>
      </div>
    </header>

    <main class="grid-container">
      
      <section class="feed-column">
        <div class="column-header">
          <div class="brand-badge ab">Aftonbladet</div>
          <span class="count" v-if="abArticles.length">{{ abArticles.length }} artiklar</span>
        </div>
        
        <div v-if="errorAB" class="status-msg error">{{ errorAB }}</div>
        
        <div class="articles-list">
          <article v-for="a in abArticles" :key="a.link" class="article-card">
            <div class="card-meta">
              <span class="date">{{ formatDate(a.published) }}</span>
            </div>
            <h3 class="title">{{ a.title }}</h3>
            <p class="summary" v-html="truncate(a.summary)"></p>
            <a :target="_blank" :href="a.link" class="read-more">Läs artikeln</a>
          </article>
        </div>
      </section>

      <section class="feed-column">
        <div class="column-header">
          <div class="brand-badge exp">Expressen</div>
          <span class="count" v-if="expArticles.length">{{ expArticles.length }} artiklar</span>
        </div>

        <div v-if="errorExp" class="status-msg error">{{ errorExp }}</div>

        <div class="articles-list">
          <article v-for="a in expArticles" :key="a.link" class="article-card">
            <div class="card-meta">
              <span class="date">{{ formatDate(a.published) }}</span>
            </div>
            <h3 class="title">{{ a.title }}</h3>
            <p class="summary" v-html="truncate(a.summary)"></p>
            <a :target="_blank" :href="a.link" class="read-more">Läs artikeln</a>
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
const loading = ref(false)
const errorAB = ref(null)
const errorExp = ref(null)

const fetchAll = async () => {
  loading.value = true
  errorAB.value = null
  errorExp.value = null
  await Promise.all([fetchAB(), fetchExp()])
  loading.value = false
}

const fetchAB = async () => {
  try {
    const res = await fetch('/api/debatt')
    const data = await res.json()
    abArticles.value = data.articles
  } catch (e) { errorAB.value = "Kunde inte hämta Aftonbladet" }
}

const fetchExp = async () => {
  try {
    const res = await fetch('/api/expressen')
    const data = await res.json()
    expArticles.value = data.articles
  } catch (e) { errorExp.value = "Kunde inte hämta Expressen" }
}

const formatDate = (d) => {
  if (!d || d === 'Saknas') return ''
  return new Date(d).toLocaleDateString('sv-SE', { day: 'numeric', month: 'short' })
}

const truncate = (text) => {
  if (!text) return ''
  return text.length > 160 ? text.substring(0, 160) + '...' : text
}

onMounted(fetchAll)
</script>

<style scoped>
/* Gemensamma variabler för ljust tema */
:root {
  --bg-color: #f4f7f9;
  --card-bg: #ffffff;
  --text-main: #1a1a1b;
  --text-muted: #5c5c5c;
  --accent: #3b82f6;
  --border: #e5e7eb;
}

.app-wrapper {
  background-color: #f4f7f9;
  min-height: 100vh;
  color: #1a1a1b;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  padding-bottom: 50px;
}

/* Header Styling */
.main-header {
  background: white;
  padding: 40px 20px;
  border-bottom: 1px solid #e5e7eb;
  text-align: center;
  margin-bottom: 30px;
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin-bottom: 8px;
}

.logo h1 {
  margin: 0;
  font-size: 2rem;
  font-weight: 800;
  letter-spacing: -0.025em;
}

.subtitle {
  color: #5c5c5c;
  margin-bottom: 24px;
}

.btn-refresh {
  background: #1a1a1b;
  color: white;
  border: none;
  padding: 10px 24px;
  border-radius: 99px;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.1s;
}

.btn-refresh:active { transform: scale(0.96); }

/* Layout Grid */
.grid-container {
  max-width: 1100px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
  padding: 0 20px;
}

@media (max-width: 800px) {
  .grid-container { grid-template-columns: 1fr; }
}

/* Column Styling */
.column-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 0 4px;
}

.brand-badge {
  font-weight: 700;
  text-transform: uppercase;
  font-size: 0.75rem;
  letter-spacing: 0.05em;
  padding: 4px 12px;
  border-radius: 4px;
}

.brand-badge.ab { background: #fee2e2; color: #b91c1c; }
.brand-badge.exp { background: #dbeafe; color: #1e40af; }

.count { font-size: 0.8rem; color: #9ca3af; }

/* Article Card */
.article-card {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 16px;
  transition: box-shadow 0.2s;
}

.article-card:hover {
  box-shadow: 0 4px 12px rgba(0,0,0,0.05);
}

.card-meta { margin-bottom: 8px; }
.date { font-size: 0.75rem; font-weight: 600; color: #9ca3af; text-transform: uppercase; }

.title {
  margin: 0 0 10px 0;
  font-size: 1.15rem;
  line-height: 1.4;
  font-weight: 700;
}

.summary {
  font-size: 0.95rem;
  color: #4b5563;
  line-height: 1.5;
  margin-bottom: 16px;
}

.read-more {
  font-size: 0.9rem;
  font-weight: 600;
  color: #3b82f6;
  text-decoration: none;
}

.read-more:hover { text-decoration: underline; }

.status-msg.error {
  background: #fff1f2;
  color: #e11d48;
  padding: 12px;
  border-radius: 8px;
  margin-bottom: 16px;
  font-size: 0.9rem;
}
</style>
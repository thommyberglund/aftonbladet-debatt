<template>
  <div :class="['app-wrapper', { 'dark-theme': isDark }]">
    <header class="main-header">
      <div class="header-content">
        <div class="top-bar">
          <button @click="toggleTheme" class="theme-toggle">
            {{ isDark ? '☀️ Ljust tema' : '🌙 Mörkt tema' }}
          </button>
        </div>
        <h1>Debattkollen</h1>
        <p class="subtitle">Senaste opinionen från Sveriges största redaktioner</p>
        <button @click="fetchAll" :disabled="loading" class="btn-refresh">
          {{ loading ? 'Uppdaterar...' : 'Uppdatera alla flöden' }}
        </button>
      </div>
    </header>

    <main class="grid-container">
      <section class="feed-column">
        <div class="column-header"><div class="brand-badge ab">Aftonbladet</div></div>
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
        <div class="column-header"><div class="brand-badge exp">Expressen</div></div>
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
        <div class="column-header"><div class="brand-badge dn">DN Debatt</div></div>
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
const isDark = ref(localStorage.getItem('theme') === 'dark')
const errorAB = ref(null), errorExp = ref(null), errorDN = ref(null)

const toggleTheme = () => {
  isDark.value = !isDark.value
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

const fetchAll = async () => {
  loading.value = true
  await Promise.all([fetchAB(), fetchExp(), fetchDN()])
  loading.value = false
}

const fetchAB = async () => {
  try {
    const res = await fetch('/api/debatt')
    const data = await res.json(); abArticles.value = data.articles
  } catch (e) { errorAB.value = "Fel vid AB" }
}

const fetchExp = async () => {
  try {
    const res = await fetch('/api/expressen')
    const data = await res.json(); expArticles.value = data.articles
  } catch (e) { errorExp.value = "Fel vid Expressen" }
}

const fetchDN = async () => {
  try {
    const res = await fetch('/api/dn')
    const data = await res.json(); dnArticles.value = data.articles
  } catch (e) { errorDN.value = "Fel vid DN" }
}

const formatDate = (d) => d && d !== 'Saknas' ? new Date(d).toLocaleDateString('sv-SE', { day: 'numeric', month: 'short' }) : ''
const truncate = (text) => text ? (text.length > 140 ? text.substring(0, 140) + '...' : text) : ''

onMounted(fetchAll)
</script>

<style>
/* CSS Variabler för teman */
:root {
  --bg-app: #f1f5f9;
  --bg-header: #ffffff;
  --bg-card: #ffffff;
  --text-title: #0f172a;
  --text-body: #334155;
  --text-meta: #64748b;
  --border-card: #cbd5e1;
  --btn-bg: #0f172a;
  --btn-text: #ffffff;
}

.dark-theme {
  --bg-app: #0f172a;
  --bg-header: #1e293b;
  --bg-card: #1e293b;
  --text-title: #f8fafc;
  --text-body: #cbd5e1;
  --text-meta: #94a3b8;
  --border-card: #334155;
  --btn-bg: #3b82f6;
  --btn-text: #ffffff;
}

.app-wrapper { 
  background: var(--bg-app); 
  min-height: 100vh; 
  font-family: 'Inter', sans-serif; 
  padding-bottom: 60px;
  color: var(--text-body);
  transition: background 0.3s, color 0.3s;
}

.main-header { 
  background: var(--bg-header); 
  padding: 20px 20px 40px; 
  border-bottom: 2px solid var(--border-card); 
  text-align: center; 
  margin-bottom: 40px; 
}

.top-bar { display: flex; justify-content: flex-end; max-width: 1400px; margin: 0 auto; }
.theme-toggle { 
  background: transparent; border: 1px solid var(--border-card); 
  color: var(--text-title); padding: 5px 12px; border-radius: 8px; cursor: pointer; font-weight: 600;
}

.main-header h1 { color: var(--text-title); font-size: 2.5rem; font-weight: 800; margin: 10px 0; }
.subtitle { color: var(--text-meta); margin-bottom: 25px; }

.grid-container { 
  max-width: 1440px; margin: 0 auto; display: grid; 
  grid-template-columns: repeat(3, 1fr); gap: 24px; padding: 0 24px; 
}

@media (max-width: 1100px) { .grid-container { grid-template-columns: repeat(2, 1fr); } }
@media (max-width: 700px) { .grid-container { grid-template-columns: 1fr; } }

.article-card { 
  background: var(--bg-card); border: 1px solid var(--border-card); 
  border-radius: 12px; padding: 24px; margin-bottom: 20px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); 
}

.title { color: var(--text-title); font-size: 1.25rem; font-weight: 800; margin: 10px 0; }
.summary { color: var(--text-body); line-height: 1.6; }
.date { font-size: 0.8rem; color: var(--text-meta); font-weight: 700; }

.btn-refresh { 
  background: var(--btn-bg); color: var(--btn-text); 
  border: none; padding: 12px 28px; border-radius: 99px; font-weight: 700; cursor: pointer; 
}

.brand-badge { font-weight: 800; padding: 6px 14px; border-radius: 6px; font-size: 0.85rem; margin-bottom: 15px; display: inline-block; }
.ab { background: #fee2e2; color: #991b1b; }
.exp { background: #dbeafe; color: #1e40af; }
.dn { background: #f1f5f9; color: #0f172a; border: 1px solid #cbd5e1; }
.dark-theme .dn { background: #334155; color: #f8fafc; border-color: #475569; }
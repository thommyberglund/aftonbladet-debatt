# Debatty: Aftonbladet Debatt Tracker 📰

Ett fullstack-projekt som samlar in och visualiserar debattartiklar från Aftonbladet i realtid. Projektet består av en Go-backend och en modern Vue.js-frontend.

## 🏗 Arkitektur

Projektet är containeriserat med Docker och använder följande struktur:

- **Frontend**: Vue 3 (Vite) serverad via Nginx.
- **Backend**: Som skrapar RSS-flöden med Feedparser.
- **Proxy**: Designad för att köras bakom Caddy eller Nginx för HTTPS-terminering.

[Image of a full stack web architecture diagram showing frontend, backend api, and a reverse proxy]

## 🚀 Snabbstart med Docker Compose

För att starta hela projektet (både frontend och backend) i en sammankopplad miljö:

1. **Bygg och starta containrarna:**
   ```bash
   docker-compose up --build -d

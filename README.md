# Debatty: Aftonbladet Debatt Tracker 📰

Ett fullstack-projekt som samlar in och visualiserar debattartiklar från Sveriges största redaktioner i realtid. Projektet består av en Go-backend och en modern Vue.js-frontend.

## 🏗 Arkitektur

Projektet är containeriserat med Docker och använder följande struktur:

- **Frontend**: Vue 3.5 (Vite 7.3) med TypeScript-stöd
- **Backend**: Go 1.25.7 som skrapar RSS-flöden med [gofeed](https://github.com/mmcdole/gofeed)
- **Proxy**: Designad för att köras bakom Caddy eller Nginx för HTTPS-terminering

```
aftonbladet-debatt/
├── backend/
│   ├── main.go       # RSS feed scraper
│   ├── go.mod        # Go module (frosteby.eu/aftonbladetdebatt)
│   └── Dockerfile    # Multi-stage build → Alpine
└── frontend/
    ├── src/
    │   ├── App.vue    # Main UI (3-column grid)
    │   ├── main.ts    # Vue entry point
    │   └── router/    # Vue Router
    └── package.json  # Vue 3, Vite, Tailwind CSS
```

## 📡 Data Sources

Följande RSS-flöden skrapas:
- **Aftonbladet Debatt**: `https://rss.aftonbladet.se/rss2/small/pages/sections/debatt/`
- **Expressen Debatt**: `https://feeds.expressen.se/debatt/`
- **DN Debatt**: `https://www.dn.se/debatt/rss/`

## 🚀 Snabbstart med Docker Compose

För att starta hela projektet (både frontend och backend) i en sammankopplad miljö:

1. **Förutsättningar**:
   - Docker och Docker Compose installerade
   - Extern nätverksresurs `caddynet` måste finnas (eller redigera `docker-compose.yaml`)

2. **Bygg och starta containrarna:**
   ```bash
   docker-compose up --build -d
   ```

3. **Åtkomst**:
   - Frontend: `http://localhost:8080`
   - Backend API: `http://localhost:8000`

## 🛠 Lokalt Utveckling

### Backend
```bash
cd backend
go mod tidy
go run main.go
```
API-endpoints:
- `GET /debatt` - Aftonbladet artiklar
- `GET /expressen` - Expressen artiklar
- `GET /dn` - DN artiklar

### Frontend
```bash
cd frontend
npm install
npm run dev
```

## 🎨 Frontend Features

- **Responsive design**: 3-kolumns layout (anpassas till 2 eller 1 kolumn på mindre skärmar)
- **Tema-växling**: Mörkt/ljust tema med lokal lagring
- **Real-time uppdatering**: Manuell uppdateringsknapp för alla flöden
- **Artikelkort**: Visar datum, titel, sammanfattning och länk till originalartikeln

## 📦 Teknisk Stack

| Komponent | Teknologi |
|-----------|-----------|
| **Backend** | Go 1.25.7, gofeed v1.3.0 |
| **Frontend** | Vue 3.5.29, TypeScript, Vite 7.3.1 |
| **Styling** | Tailwind CSS 4.2.1, CSS Variables |
| **Routing** | Vue Router 5.0.3 |
| **Container** | Docker, Alpine Linux |

## 🔧 Konfiguration

### Miljövariabler (Backend)

| Variabel | Standard | Beskrivning |
|----------|----------|--------------|
| `PORT` | 8000 | Serverport |

### Miljövariabler (Frontend)

| Variabel | Standard | Beskrivning |
|----------|----------|--------------|
| `VITE_API_BASE_URL` | `/api` | Bas-URL för API-anrop |

## 🤝 Bidra

1. Forka repositoryt
2. Skapa en feature branch (`git checkout -b feature/ny-funktion`)
3. Committa dina ändringar (`git commit -m 'Lägg till ny funktion'`)
4. Pusha till branch (`git push origin feature/ny-funktion`)
5. Öppna en Pull Request

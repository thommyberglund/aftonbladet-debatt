# Debatty Backend (Go)

Snabb och resurssnål service byggd i Go för att skrapa och leverera debattartiklar från svenska nyhetssajter via RSS-flöden.

## 📡 API Endpoints

| Endpoint | Metod | Beskrivning | RSS Källa |
|----------|-------|--------------|-----------|
| `/debatt` | GET | Hämta Aftonbladet Debatt artiklar | `https://rss.aftonbladet.se/rss2/small/pages/sections/debatt/` |
| `/expressen` | GET | Hämta Expressen Debatt artiklar | `https://feeds.expressen.se/debatt/` |
| `/dn` | GET | Hämta DN Debatt artiklar | `https://www.dn.se/debatt/rss/` |

**Svarformat** (JSON):
```json
{
  "articles": [
    {
      "title": "Artikel titel",
      "published": "2024-01-01T00:00:00+01:00",
      "link": "https://example.com/artikel",
      "summary": "Artikel sammanfattning..."
    }
  ]
}
```

## 🏗 Teknisk Stack

- **Runtime**: Go 1.25.7
- **RSS Parser**: [gofeed v1.3.0](https://github.com/mmcdole/gofeed)
- **HTTP Server**: Standard `net/http`
- **CORS**: Aktiverat för alla ursprung (`Access-Control-Allow-Origin: *`)

## 🛠 Installation (Lokalt)

1. **Förutsättningar**:
   - Go 1.21+ installerat

2. **Installera beroenden**:
   ```bash
   go mod tidy
   ```

3. **Kör servern**:
   ```bash
   go run main.go
   ```
   Servern startar på `http://localhost:8000`

## 🐳 Docker

### Bygg image
```bash
cd backend
docker build -t debatt-api-go .
```

### Kör container
```bash
docker run -p 8000:8000 debatt-api-go
```

### Docker Compose (hela projektet)
Se root `docker-compose.yaml` för att köra backend + frontend tillsammans.

## 📁 Projektstruktur

```
backend/
├── main.go       # Huvudapplikation med HTTP handlers
├── go.mod        # Go module (frosteby.eu/aftonbladetdebatt)
├── go.sum        # Checksum för beroenden
└── Dockerfile    # Multi-stage build → Alpine Linux
```

## 🔧 Miljövariabler

| Variabel | Standard | Beskrivning |
|----------|----------|--------------|
| `PORT` | 8000 | Serverport (ändra i `main.go:68`)

## 📝 Kodstruktur

### `main.go`

**Structs**:
- `Article`: Representerar en artikel med fält för titel, publiceringsdatum, länk och sammanfattning

**Funktioner**:
- `fetchFeed(url string)`: Skrapar RSS-feed från given URL
- `getAftonbladet()`, `getExpressen()`, `getDN()`: HTTP handlers för respektive källa
- `setupResponse()`: Konfigurerar CORS och Content-Type headers
- `handleResult()`: Hanterar JSON-svar och fel

## 🔒 Säkerhet

⚠️ **CORS**: Backend tillåter för närvarande alla ursprung (`*`). För produktion:
- Begränsa till specifika domäner
- Lägg till autentiseringsmekanism om nödvändigt

## 🧪 Testning

Inga enhetstester finns för närvarande. Rekommenderade tester:
- Test `fetchFeed()` med mockade RSS-svar
- Test HTTP handlers och statuskoder
- Test JSON-serialisering av `Article` struct

## 🤝 Bidra

1. Forka repositoryt
2. Skapa en feature branch
3. Gör dina ändringar
4. Skriv tester (om tillämpligt)
5. Öppna en Pull Request

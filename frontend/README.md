# Aftonbladet Debatt - Frontend

Vue 3 webbgränssnitt för att visa debattartiklar.

## Tekniker
- Vue 3 (Vite)
- Tailwind CSS

## Installation (Lokalt)
1. Installera paket:
   npm install

2. Kör i utvecklingsläge:
   npm run dev

3. Bygg för produktion:
   npm run build

## Nätverkskonfiguration
Appen är inställd på att anropa `/api/debatt`. 
I produktion hanteras detta via en Reverse Proxy (Caddy/Nginx) som mappar `/api/*` till backend-containern på port 8000.

## Docker
Bygg image:
docker build -t debatt-ui .

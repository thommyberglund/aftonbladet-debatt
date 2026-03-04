# Aftonbladet Debatt - Backend

Python-tjänst som hämtar debattartiklar via RSS och serverar dem som JSON.

## Tekniker
- FastAPI
- Feedparser
- Uvicorn

## Installation (Lokalt)
1. Installera beroenden:
   pip install -r requirements.txt

2. Starta servern:
   uvicorn main:app --reload --port 8000

## API Endpoints
- GET `/debatt` : Returnerar lista på artiklar.
- GET `/docs`   : Interaktiv dokumentation (Swagger).

## Docker
Bygg image:
docker build -t debatt-api .

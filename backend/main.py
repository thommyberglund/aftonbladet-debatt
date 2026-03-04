from fastapi import FastAPI, HTTPException
import feedparser
from starlette.middleware.cors import CORSMiddleware

app = FastAPI(title="Aftonbladet Debatt API")

# Konfigurera User-Agent globalt för feedparser
feedparser.USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

# TILLÅT FRONTEND ATT ANROPA BACKEND
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"], # I produktion, sätt din specifika URL här
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/debatt")
async def get_debates():
    rss_url = "https://rss.aftonbladet.se/rss2/small/pages/sections/debatt/"

    feed = feedparser.parse(rss_url)

    # Kontrollera om något gick snett vid hämtningen
    if feed.bozo != 0:
        raise HTTPException(status_code=500, detail=f"Kunde inte hämta flödet: {feed.bozo_exception}")

    results = []

    for entry in feed.entries:
        results.append({
            "title": entry.get('title', 'Information saknas'),
            "published": entry.get('published', 'Information saknas'),
            "link": entry.get('link', 'Information saknas'),
            "summary": entry.get('summary', 'Information saknas')
        })

    return {
        "count": len(results),
        "articles": results
    }


if __name__ == "__main__":
    import uvicorn

    # Kör servern på localhost:8000
    uvicorn.run(app, host="0.0.0.0", port=8000)
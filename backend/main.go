package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mmcdole/gofeed"
)

// Article motsvarar den data vi skickar till frontenden
type Article struct {
	Title     string `json:"title"`
	Published string `json:"published"`
	Link      string `json:"link"`
	Summary   string `json:"summary"`
}

func getDebates(w http.ResponseWriter, r *http.Request) {
	// Tillåt CORS (viktigt för frontend)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fp := gofeed.NewParser()
	// Sätt en User-Agent för att undvika blockering
	fp.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

	feed, err := fp.ParseURL("https://rss.aftonbladet.se/rss2/small/pages/sections/debatt/")
	if err != nil {
		http.Error(w, "Kunde inte hämta RSS: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var articles []Article
	for _, item := range feed.Items {
		articles = append(articles, Article{
			Title:     item.Title,
			Published: item.Published,
			Link:      item.Link,
			Summary:   item.Description,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"articles": articles,
	})
}

func main() {
	http.HandleFunc("/debatt", getDebates)

	log.Println("Backend körs på port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mmcdole/gofeed"
)

type Article struct {
	Title     string `json:"title"`
	Published string `json:"published"`
	Link      string `json:"link"`
	Summary   string `json:"summary"`
}

func fetchFeed(url string) ([]Article, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
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
	return articles, nil
}

// Handlers för respektive källa
func getAftonbladet(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	articles, err := fetchFeed("https://rss.aftonbladet.se/rss2/small/pages/sections/debatt/")
	handleResult(w, articles, err)
}

func getExpressen(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	articles, err := fetchFeed("https://feeds.expressen.se/debatt/")
	handleResult(w, articles, err)
}

func getDN(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	articles, err := fetchFeed("https://www.dn.se/debatt/rss/")
	handleResult(w, articles, err)
}

// Hjälpfunktioner för HTTP-svar
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}

func handleResult(w http.ResponseWriter, articles []Article, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"articles": articles})
}

func main() {
	http.HandleFunc("/debatt", getAftonbladet)
	http.HandleFunc("/expressen", getExpressen)
	http.HandleFunc("/dn", getDN)

	log.Println("Go Backend körs på port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
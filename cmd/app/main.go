package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"daily-news/pkg/newsdata"
	"daily-news/pkg/siteserver"
	"daily-news/site"
)

type newsResponse struct {
	Articles []newsdata.ArticleEntry `json:"articles"`
}

func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/apis/v1/health", healthHandler)
	apiMux.HandleFunc("/apis/v1/news", newsHandler)

	handler, err := siteserver.WrapHandler(apiMux, site.DistDirFS)
	if err != nil {
		log.Fatalf("failed to init site handler: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("daily-news starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"source": "go-embedded-spa",
	})
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if !newsdata.IsValidCategory(category) {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid category",
		})
		return
	}

	language := newsdata.NormalizeLanguage(r.URL.Query().Get("language"))
	articles, err := newsdata.LoadCategory(newsdata.Category(category), language)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to load data",
		})
		return
	}

	writeJSON(w, http.StatusOK, newsResponse{
		Articles: articles,
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"error":"encoding_failed"}`, http.StatusInternalServerError)
	}
}

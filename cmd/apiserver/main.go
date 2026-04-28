package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"daily-news/pkg/newsdata"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", healthHandler)
	mux.HandleFunc("/api/v1/ai", aiDailyHandler)

	handler := withCORS(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "17632"
	}

	log.Printf("daily-news api listening on :%s (data_root=%s)", port, debugDataRoot())
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func debugDataRoot() string {
	root, err := newsdata.DataRoot()
	if err != nil {
		return "(unavailable)"
	}
	return root
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"service": "daily-news-api",
	})
}

func aiDailyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, http.StatusMethodNotAllowed, "method_not_allowed", "仅支持 GET")
		return
	}

	date := r.URL.Query().Get("date")
	raw, err := newsdata.LoadRawDailyJSON(newsdata.CategoryAI, date)
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
		return
	}

	switch {
	case errors.Is(err, newsdata.ErrMissingDateParam):
		writeErr(w, http.StatusBadRequest, "missing_date", "缺少查询参数 date，格式为 YYYY-MM-DD，例如 ?date=2026-04-28")
	case errors.Is(err, newsdata.ErrInvalidDateParam):
		writeErr(w, http.StatusBadRequest, "invalid_date", "date 格式无效，应为 YYYY-MM-DD")
	case errors.Is(err, newsdata.ErrDailyFileNotFound):
		writeErr(w, http.StatusNotFound, "not_found", "当日 ai 栏目数据不存在，请确认已生成 ~/.daily-news/data/ai/"+date+".json")
	case errors.Is(err, newsdata.ErrDailyFileNotJSON):
		writeErr(w, http.StatusInternalServerError, "invalid_file", "磁盘上的当日 JSON 无法解析，请检查数据文件内容")
	default:
		log.Printf("api /api/v1/ai: %v", err)
		writeErr(w, http.StatusInternalServerError, "internal_error", "读取数据失败")
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type errBody struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func writeErr(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, errBody{Error: code, Message: message})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"error":"encoding_failed"}`, http.StatusInternalServerError)
	}
}

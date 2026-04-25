package siteserver

import (
	"bytes"
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"
)

func WrapHandler(apiHandler http.Handler, staticFS fs.FS) (http.Handler, error) {
	indexHTML, err := fs.ReadFile(staticFS, "index.html")
	if err != nil {
		return nil, err
	}

	fileServer := http.FileServer(http.FS(staticFS))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/apis/") {
			apiHandler.ServeHTTP(w, r)
			return
		}

		requestPath := strings.TrimPrefix(r.URL.Path, "/")
		if requestPath == "" {
			requestPath = "index.html"
		}

		if hasExtension(requestPath) {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			fileServer.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Cache-Control", "no-cache")
		http.ServeContent(w, r, "index.html", time.Time{}, bytes.NewReader(indexHTML))
	}), nil
}

func hasExtension(p string) bool {
	ext := path.Ext(p)
	return ext != ""
}

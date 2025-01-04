package handlers

import (
	"net/http"

	indexPage "github.com/ropfoo/gothulhu/internal/tmpl/pages/index"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "public, max-age=86400") // Cache for 24 hours
	w.Write([]byte(indexPage.IndexPage()))
}

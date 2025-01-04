package handlers

import (
	"net/http"

	apiPage "github.com/ropfoo/gothulhu/internal/tmpl/pages/api"
)

func ApiDocsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(apiPage.ApiPage()))
}

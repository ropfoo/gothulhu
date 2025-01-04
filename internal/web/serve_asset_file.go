package web

import (
	"net/http"
)

func ServeAssetFile(w http.ResponseWriter, r *http.Request, path string) {

	http.ServeFile(w, r, "/assets/"+path)
}

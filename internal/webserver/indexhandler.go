package webserver

import (
	"net/http"
	"path/filepath"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// handle files
	path := string(config.HTTP_BASE_DIR) + r.URL.Path

	// check if it's a file request
	if filepath.Ext(path) != "" {
		http.FileServer(config.HTTP_BASE_DIR).ServeHTTP(w, r)
		return
	}

	// if request has path, then its already a 404
	if r.URL.Path != "/" {
		w.Write(NotFound(""))
		return
	}

	// then handle index
	var activated, errCode = internal.CheckActivation()

	if activated {
		ServeHTML(w, "deactivate")
		return
	}

	if errCode == config.CHECK_ACTIVATION_READFAILURE {
		ServeHTML(w, "activate")
		return
	}

	if errCode == config.CHECK_ACTIVATION_NOTACTIVATED {
		ServeHTML(w, "reactivate")
		return
	}
}

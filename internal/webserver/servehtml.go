package webserver

import (
	"net/http"
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func ServeHTML(w http.ResponseWriter, html string) {
	f, err := os.ReadFile(string(config.HTTP_BASE_DIR) + html + ".html")
	if err != nil {
		f = NotFound(err.Error())
	}

	w.Write([]byte(f))
}

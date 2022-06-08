package webserver

import (
	"net/http"
	"os"
	"strings"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func ServeHTML(w http.ResponseWriter, html string) {
	f, err := os.ReadFile(string(config.HTTP_BASE_DIR) + html + ".html")
	if err != nil {
		f = NotFound(err.Error())
	}

	// add system info
	toString := string(f)

	finalString := strings.Replace(
		toString,
		"<!--#####SYSTEM_INFO#####-->",
		"<strong>Computer Name:</strong> "+config.ComputerName,
		-1,
	)

	w.Write([]byte(finalString))
}

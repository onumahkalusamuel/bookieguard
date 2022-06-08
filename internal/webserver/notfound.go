package webserver

import (
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func NotFound(msg string) []byte {

	var f []byte

	var err error

	f, err = os.ReadFile(string(config.HTTP_BASE_DIR) + "404.html")
	if err != nil {
		f = []byte("404: Not found.")
	}

	// pass in the message
	if len(f) < 1 {
		f = []byte(msg)
	}

	return f
}

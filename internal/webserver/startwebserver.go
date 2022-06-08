package webserver

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func StartWebServer() {

	// index and catch all
	http.HandleFunc("/", IndexHandler)

	// others
	http.HandleFunc("/deactivate", DeactivateHandler)

	http.HandleFunc("/activate", ActivateHandler)

	fmt.Printf("Starting server at port %v\n", config.WEB_PORT)
	if err := http.ListenAndServe(net.JoinHostPort(config.WEB_HOST, config.WEB_PORT), nil); err != nil {
		log.Fatal(err)
	}
}

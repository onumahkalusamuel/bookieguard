package webserver

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/itrepablik/itrlog"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
)

func StartWebServer() {

	go func() {

		time.Sleep(3 * time.Second)

		var activated, _ = internal.CheckActivation()

		// launch server
		if activated {
			fmt.Println("Already activated.")
			internal.StartServer()
			return
		}

		// launch browser
		if !activated {
			LaunchBrowser()
		}
	}()

	// index and catch all
	http.HandleFunc("/", IndexHandler)

	// others
	http.HandleFunc("/deactivate", DeactivateHandler)

	http.HandleFunc("/activate", ActivateHandler)

	fmt.Printf("Starting server at port %v\n", config.WEB_PORT)
	if err := http.ListenAndServe(net.JoinHostPort(config.WEB_HOST, config.WEB_PORT), nil); err != nil {
		itrlog.Error(err)
	}

	fmt.Println("Did it get here?")
}

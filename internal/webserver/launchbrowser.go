package webserver

import (
	"fmt"
	"net"
	"os/exec"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func LaunchBrowser() {

	fmt.Println("Opening Admin Panel in browser...")

	err := exec.Command(
		"rundll32",
		"url.dll,FileProtocolHandler",
		"http://"+net.JoinHostPort(config.WEB_HOST, config.WEB_PORT),
	).Start()

	if err != nil {
		fmt.Println("An Error occured: ", err)
	}
}

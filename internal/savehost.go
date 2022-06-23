package internal

import (
	"os"
	"strings"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func SaveHost(host string) {

	var exists bool
	var filecontent []byte
	var err error

	// load the host file
	filecontent, _ = os.ReadFile(config.HostsFile)

	hosts := string(filecontent)

	// open or create host file
	g, err := os.OpenFile(config.HostsFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {

	}
	defer g.Close()

	if strings.Contains(hosts, host) {
		exists = true
	}

	if !exists {
		g.WriteString(host + ",")
	}
}

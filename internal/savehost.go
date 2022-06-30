package internal

import (
	"os"
	"strings"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func SaveHost(host string) {

	host = strings.Split(host, ":")[0]
	host = strings.Trim(host, "www.")

	// if host lent is more than three pick, the last three
	hostparts := strings.Split(host, ".")
	length := len(hostparts)
	if length > 3 {
		host = hostparts[length-1] + hostparts[length-2] + hostparts[length-3]
	}

	// its more than 4 just skip
	if length > 4 {
		return
	}

	// if len(hostparts)

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

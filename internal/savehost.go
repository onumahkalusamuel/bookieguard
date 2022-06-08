package internal

import (
	"fmt"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func SaveHost(host string) {

	var exists bool

	for _, h := range config.Hosts {
		if host == h {
			exists = true
			break
		}
	}
	if !exists {
		config.Hosts = append(config.Hosts, host)
	}

	fmt.Println(config.Hosts)
}

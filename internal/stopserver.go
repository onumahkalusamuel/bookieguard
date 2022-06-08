package internal

import (
	"github.com/onumahkalusamuel/bookieguard/config"
)

func StopServer() (bool, error) {

	err := config.PROXY_SERVER_HANDLE.Shutdown(nil)
	if err != nil {
		return false, err
	}

	return true, nil

}

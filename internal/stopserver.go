package internal

import (
	"context"
	"net/http"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func StopServer() (bool, error) {

	err := config.PROXY_SERVER_HANDLE.Shutdown(context.TODO())
	if err != nil {
		return false, err
	}

	config.PROXY_SERVER_HANDLE = new(http.Server)

	// go ResetProxy()

	return true, nil

}

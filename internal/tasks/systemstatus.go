package tasks

import (
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
)

func SystemStatus() {

	active, _ := internal.CheckActivation()

	if active {
		json_data := config.BodyStructure{}
		internal.PostRequest(config.Endpoints["system-status"], json_data)
	}
}

package tasks

import (
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
)

func SendHosts() {
	hosts, _ := os.ReadFile(config.HostsFile)
	if len(string(hosts)) < 5 {
		return
	}
	json_data := config.BodyStructure{
		"hosts": string(hosts),
	}
	internal.PostRequest(config.Endpoints["upload-hosts"], json_data)
}

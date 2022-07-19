package tasks

import (
	"fmt"
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
	// empty the hosts file
	f, err := os.Create(config.HostsFile)
	if err != nil {
		fmt.Println(err)
	}
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}

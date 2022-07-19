package webserver

import (
	"net/http"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

// this handler will check for updates, fix all needed,
// and tell user to download new app version is available.
func UpdatesHandler(w http.ResponseWriter, r *http.Request) {

	blocklist := internal.FetchBlockList()
	blHash := pkg.GetHash(blocklist)
	json_data := config.BodyStructure{
		"blocklistHash": blHash,
		"appVersion":    config.AppVersion,
	}
	resp := internal.PostRequest(config.Endpoints["update"], json_data)

	if resp["success"] == "false" {
		ServeJSON(w, resp)
		return
	}

	if resp["expired"] == "true" {
		internal.SaveBlocklist("*")
		internal.StopServer()
		ServeJSON(w, config.BodyStructure{
			"success":  "false",
			"message":  "Subscription expired. Please resubscribe",
			"redirect": config.WebBase,
		})
		return
	}

	// blocklist update
	if resp["blocklistHash"] != blHash {
		// update blocklist
		internal.SaveBlocklist(resp["blocklist"])
	}

	if resp["appVersion"] != "" && resp["appVersion"] != config.AppVersion {
		ServeJSON(w, config.BodyStructure{
			"success":  "true",
			"message":  "New program update is available. Please download and install",
			"redirect": config.WebBase,
		})
		return
	}

	ServeJSON(w, config.BodyStructure{
		"success": "true",
		"message": "System is up to date.",
	})

}

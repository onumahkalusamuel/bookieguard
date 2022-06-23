package tasks

import (
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func Update() {
	blocklist := internal.FetchBlockList()
	blHash := pkg.GetHash(blocklist)
	json_data := config.BodyStructure{
		"blocklistHash": blHash,
	}
	resp := internal.PostRequest(config.Endpoints["update"], json_data)

	if resp["success"] == "false" {
		return
	}

	if resp["expired"] == "true" {
		internal.SaveBlocklist("*")
		internal.StopServer()
		return
	}

	// blocklist update
	if resp["blocklistHash"] != blHash {
		// update blocklist
		internal.SaveBlocklist(resp["blocklist"])
	}
}

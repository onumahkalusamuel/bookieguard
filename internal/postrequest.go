package internal

import (
	"net/http"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func PostRequest(url string, json_data config.BodyStructure) config.BodyStructure {

	resp, err := http.Post(url, "application/json", MarshalRequest(json_data))

	if err != nil {
		return config.BodyStructure{
			"success": "false",
			"message": "Network error",
		}
	}

	output, err := UnmarshalResponse(resp)

	if err != nil {
		return config.BodyStructure{
			"success": "false",
			"message": "Response not understood. please try again later",
		}
	}

	return output
}

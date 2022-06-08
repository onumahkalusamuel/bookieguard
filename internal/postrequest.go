package internal

import (
	"log"
	"net/http"
)

func PostRequest(url string, json_data map[string]string) map[string]string {

	resp, err := http.Post(url, "application/json", MarshalRequest(json_data))

	if err != nil {
		log.Fatal(err)
	}

	output, err := UnmarshalResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

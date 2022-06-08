package webserver

import (
	"net/http"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
)

func ActivateHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		ServeJSON(w, config.BodyStructure{
			"success": "false",
			"message": "Unable to process form fields. Please check the input and try again.",
		})
		return
	}

	config.Email = r.FormValue("email")
	config.Shop = r.FormValue("shop")

	json_data := config.BodyStructure{
		"email":        config.Email,
		"shop":         config.Shop,
		"computerName": config.ComputerName,
	}

	output := internal.PostRequest(config.Endpoints["activation"], json_data)

	if output["success"] == "false" {
		ServeJSON(w, config.BodyStructure{
			"success": "false",
			"message": output["message"],
		})
		return
	}

	if output["success"] == "true" && output["activated"] == "true" {
		activated, err := internal.Activate(output)
		if err != nil {
			ServeJSON(w, config.BodyStructure{
				"success": "false",
				"message": err.Error(),
			})
			return
		}

		if activated {
			ServeJSON(w, config.BodyStructure{
				"success":  "true",
				"message":  "Service activated successfully",
				"redirect": "/",
			})
			return
		}
	}

	ServeJSON(w, config.BodyStructure{
		"success": "false",
		"message": "Please try again later or contact Bookie Guard support.",
	})

}

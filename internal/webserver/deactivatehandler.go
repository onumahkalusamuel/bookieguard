package webserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func DeactivateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ServeJSON(w, config.BodyStructure{
			"success": "false",
			"message": "Unable to process form fields. Please check the input and try again.",
		})
		return
	}

	config.Email = r.FormValue("email")
	config.UnlockCode = r.FormValue("unlockcode")

	// vet submitted data
	if !pkg.ValidateEmail(config.Email) || config.UnlockCode == "" {
		ServeJSON(w, config.BodyStructure{
			"success": "false",
			"message": "Please fill the form completely.",
		})
		return
	}

	// check if entered email and code match
	loaded, content := internal.LoadActivationFile()
	if !loaded {
		ServeJSON(w, config.BodyStructure{
			"success": "false",
			"message": "An error occured while trying to verify credentials. Please contact admin.",
		})
		return
	}

	fmt.Println(content)

	if content["email"] != config.Email || content["unlockCode"] != config.UnlockCode {
		ServeJSON(w, config.BodyStructure{
			"success": "false",
			"message": "Wrong credentials provided!!!",
		})
		return
	}

	_, err := internal.StopServer()

	if err != nil {
		ServeJSON(w, config.BodyStructure{
			"success": "false",
			"message": "Unable to deactivate server at the moment. Please try again later.",
		})
		return
	}

	// remove files
	os.Remove(config.ActivationFile)
	os.Remove(config.BlocklistFile)

	// return response
	ServeJSON(w, config.BodyStructure{
		"success":  "true",
		"message":  "Service deactivated successfully. Don't forget to reset your proxy settings.",
		"redirect": "/",
	})

}

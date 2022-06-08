package internal

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/onumahkalusamuel/bookieguard/config"
)

func StartActivation() {

	GetUserInfo()

	json_data := map[string]string{
		"email":        config.Email,
		"shop":         config.Shop,
		"computerName": config.ComputerName,
	}

	output := PostRequest(config.Endpoints["activation"], json_data)

	if output["success"] == "false" {
		ct.Foreground(ct.Red, false)
		fmt.Printf("%v\n", output["message"])
		ct.ResetColor()
		fmt.Print("Exiting application now...\n")
		return
	}

	if output["success"] == "true" && output["activated"] == "true" {
		Activate(output)
		return
	}

	fmt.Println("Unable to activate. Please try again later.")
}

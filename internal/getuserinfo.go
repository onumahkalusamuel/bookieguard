package internal

import (
	"fmt"
	"os"
	"time"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/manifoldco/promptui"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/pkg"
)

func GetUserInfo() {

	fmt.Println("Welcome to Bookie Guard Activation Panel.")
	fmt.Println("In order to get you up and running, we need some details.")
	fmt.Println()

	for {

		prompt := promptui.Prompt{Label: "Please enter the email used in making payment"}

		result, err := prompt.Run()
		if err != nil {
			os.Exit(0)
		}

		if !pkg.ValidateEmail(result) {
			ct.Foreground(ct.Red, false)
			fmt.Println("Please enter a valid email address.")
			ct.ResetColor()
			continue
		}

		config.Email = result
		break
	}

	for {

		prompt := promptui.Select{
			Label: "Please select your shop type:",
			Items: config.Shops,
		}

		_, result, err := prompt.Run()
		if err != nil {
			os.Exit(0)
		}

		if err != nil {
			ct.Foreground(ct.Red, false)
			fmt.Println("Please select a valid shop number.")
			ct.ResetColor()
			continue
		}

		config.Shop = result
		break
	}

	time.Sleep(1 * time.Second)
}

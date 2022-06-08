package cmd

import (
	"fmt"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up Bookie Guard on this system",
	Long:  "Use this command to activate or deactivate Bookie Guard.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		var activated, errCode = internal.CheckActivation()

		if activated {
			fmt.Println("Bookie Guard already activated.")
			fmt.Println("To deactivate, complete the form below...")
			internal.Deactivate()
			return
		}

		if errCode == config.CHECK_ACTIVATION_READFAILURE {
			fmt.Println("To start using the system, please follow the steps below...")
			internal.StartActivation()
			return
		}

		if errCode == config.CHECK_ACTIVATION_NOTACTIVATED {
			fmt.Println("Your copy of Bookie Guard has expired.")
			fmt.Println("Complete the online activation then restart your system for reactivation.")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

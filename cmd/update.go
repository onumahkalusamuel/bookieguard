package cmd

import (
	"github.com/onumahkalusamuel/bookieguard/internal"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Bookie Guard",
	Long:  "Check for and pull in updates of the software and settings.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

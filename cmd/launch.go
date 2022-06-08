package cmd

import (
	"fmt"
	"net"
	"os/exec"
	"time"

	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/spf13/cobra"
)

var lauchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch Admin Panel",
	Long:  "Open Bookie Guard admin panel in your browser.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		Launch()
	},
}

func init() {
	rootCmd.AddCommand(lauchCmd)
}

func Launch() {

	fmt.Println("Opening Admin Panel in browser...")

	err := exec.Command(
		"rundll32",
		"url.dll,FileProtocolHandler",
		"http://"+net.JoinHostPort(config.WEB_HOST, config.WEB_PORT),
	).Start()

	if err != nil {
		fmt.Println("")
	}

	time.Sleep(10 * time.Second)
}
